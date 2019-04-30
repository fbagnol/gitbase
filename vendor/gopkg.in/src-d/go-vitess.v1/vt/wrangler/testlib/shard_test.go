/*
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreedto in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testlib

import (
	"strings"
	"testing"

	"golang.org/x/net/context"

	"gopkg.in/src-d/go-vitess.v1/vt/logutil"
	"gopkg.in/src-d/go-vitess.v1/vt/topo"
	"gopkg.in/src-d/go-vitess.v1/vt/topo/memorytopo"
	"gopkg.in/src-d/go-vitess.v1/vt/topotools"
	"gopkg.in/src-d/go-vitess.v1/vt/vttablet/tmclient"
	"gopkg.in/src-d/go-vitess.v1/vt/wrangler"

	topodatapb "gopkg.in/src-d/go-vitess.v1/vt/proto/topodata"
)

func TestDeleteShardCleanup(t *testing.T) {
	ctx := context.Background()
	ts := memorytopo.NewServer("cell1", "cell2")
	wr := wrangler.New(logutil.NewConsoleLogger(), ts, tmclient.NewTabletManagerClient())
	vp := NewVtctlPipe(t, ts)
	defer vp.Close()

	// Create a master, a couple good slaves
	master := NewFakeTablet(t, wr, "cell1", 0, topodatapb.TabletType_MASTER, nil)
	slave := NewFakeTablet(t, wr, "cell1", 1, topodatapb.TabletType_REPLICA, nil)
	remoteSlave := NewFakeTablet(t, wr, "cell2", 2, topodatapb.TabletType_REPLICA, nil)

	// Build keyspace graph
	err := topotools.RebuildKeyspace(context.Background(), logutil.NewConsoleLogger(), ts, master.Tablet.Keyspace, []string{"cell1", "cell2"})
	if err != nil {
		t.Fatalf("RebuildKeyspaceLocked failed: %v", err)
	}

	// Delete the ShardReplication record in cell2
	if err := ts.DeleteShardReplication(ctx, "cell2", remoteSlave.Tablet.Keyspace, remoteSlave.Tablet.Shard); err != nil {
		t.Fatalf("DeleteShardReplication failed: %v", err)
	}

	// Now try to delete the shard without even_if_serving or
	// recursive flag, should fail on serving check first.
	if err := vp.Run([]string{
		"DeleteShard",
		master.Tablet.Keyspace + "/" + master.Tablet.Shard,
	}); err == nil || !strings.Contains(err.Error(), "is still serving, cannot delete it") {
		t.Fatalf("DeleteShard() returned wrong error: %v", err)
	}

	// Now try to delete the shard with even_if_serving, but
	// without recursive flag, should fail on existing tablets.
	if err := vp.Run([]string{
		"DeleteShard",
		"-even_if_serving",
		master.Tablet.Keyspace + "/" + master.Tablet.Shard,
	}); err == nil || !strings.Contains(err.Error(), "use -recursive or remove them manually") {
		t.Fatalf("DeleteShard(evenIfServing=true) returned wrong error: %v", err)
	}

	// Now try to delete the shard with even_if_serving and recursive,
	// it should just work.
	if err := vp.Run([]string{
		"DeleteShard",
		"-recursive",
		"-even_if_serving",
		master.Tablet.Keyspace + "/" + master.Tablet.Shard,
	}); err != nil {
		t.Fatalf("DeleteShard(recursive=true, evenIfServing=true) should have worked but returned: %v", err)
	}

	// Make sure all tablets are gone.
	for _, ft := range []*FakeTablet{master, slave, remoteSlave} {
		if _, err := ts.GetTablet(ctx, ft.Tablet.Alias); !topo.IsErrType(err, topo.NoNode) {
			t.Errorf("tablet %v is still in topo: %v", ft.Tablet.Alias, err)
		}
	}

	// Make sure the shard is gone.
	if _, err := ts.GetShard(ctx, master.Tablet.Keyspace, master.Tablet.Shard); !topo.IsErrType(err, topo.NoNode) {
		t.Errorf("shard %v/%v is still in topo: %v", master.Tablet.Keyspace, master.Tablet.Shard, err)
	}
}
