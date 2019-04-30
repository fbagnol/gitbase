// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vtworkerservice.proto

package vtworkerservice // import "gopkg.in/src-d/go-vitess.v1/vt/proto/vtworkerservice"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import vtworkerdata "gopkg.in/src-d/go-vitess.v1/vt/proto/vtworkerdata"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VtworkerClient is the client API for Vtworker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VtworkerClient interface {
	// ExecuteVtworkerCommand allows to run a vtworker command by specifying the
	// same arguments as on the command line.
	ExecuteVtworkerCommand(ctx context.Context, in *vtworkerdata.ExecuteVtworkerCommandRequest, opts ...grpc.CallOption) (Vtworker_ExecuteVtworkerCommandClient, error)
}

type vtworkerClient struct {
	cc *grpc.ClientConn
}

func NewVtworkerClient(cc *grpc.ClientConn) VtworkerClient {
	return &vtworkerClient{cc}
}

func (c *vtworkerClient) ExecuteVtworkerCommand(ctx context.Context, in *vtworkerdata.ExecuteVtworkerCommandRequest, opts ...grpc.CallOption) (Vtworker_ExecuteVtworkerCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Vtworker_serviceDesc.Streams[0], "/vtworkerservice.Vtworker/ExecuteVtworkerCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &vtworkerExecuteVtworkerCommandClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Vtworker_ExecuteVtworkerCommandClient interface {
	Recv() (*vtworkerdata.ExecuteVtworkerCommandResponse, error)
	grpc.ClientStream
}

type vtworkerExecuteVtworkerCommandClient struct {
	grpc.ClientStream
}

func (x *vtworkerExecuteVtworkerCommandClient) Recv() (*vtworkerdata.ExecuteVtworkerCommandResponse, error) {
	m := new(vtworkerdata.ExecuteVtworkerCommandResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VtworkerServer is the server API for Vtworker service.
type VtworkerServer interface {
	// ExecuteVtworkerCommand allows to run a vtworker command by specifying the
	// same arguments as on the command line.
	ExecuteVtworkerCommand(*vtworkerdata.ExecuteVtworkerCommandRequest, Vtworker_ExecuteVtworkerCommandServer) error
}

func RegisterVtworkerServer(s *grpc.Server, srv VtworkerServer) {
	s.RegisterService(&_Vtworker_serviceDesc, srv)
}

func _Vtworker_ExecuteVtworkerCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(vtworkerdata.ExecuteVtworkerCommandRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VtworkerServer).ExecuteVtworkerCommand(m, &vtworkerExecuteVtworkerCommandServer{stream})
}

type Vtworker_ExecuteVtworkerCommandServer interface {
	Send(*vtworkerdata.ExecuteVtworkerCommandResponse) error
	grpc.ServerStream
}

type vtworkerExecuteVtworkerCommandServer struct {
	grpc.ServerStream
}

func (x *vtworkerExecuteVtworkerCommandServer) Send(m *vtworkerdata.ExecuteVtworkerCommandResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Vtworker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vtworkerservice.Vtworker",
	HandlerType: (*VtworkerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecuteVtworkerCommand",
			Handler:       _Vtworker_ExecuteVtworkerCommand_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "vtworkerservice.proto",
}

func init() {
	proto.RegisterFile("vtworkerservice.proto", fileDescriptor_vtworkerservice_4efa3310356e3c00)
}

var fileDescriptor_vtworkerservice_4efa3310356e3c00 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2b, 0x29, 0xcf,
	0x2f, 0xca, 0x4e, 0x2d, 0x2a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x47, 0x13, 0x96, 0x12, 0x82, 0x09, 0xa4, 0x24, 0x96, 0x24, 0x42, 0x14, 0x19,
	0x35, 0x33, 0x72, 0x71, 0x84, 0x41, 0x85, 0x85, 0xca, 0xb9, 0xc4, 0x5c, 0x2b, 0x52, 0x93, 0x4b,
	0x4b, 0x52, 0x61, 0x42, 0xce, 0xf9, 0xb9, 0xb9, 0x89, 0x79, 0x29, 0x42, 0xda, 0x7a, 0x28, 0x7a,
	0xb1, 0xab, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x91, 0xd2, 0x21, 0x4e, 0x71, 0x71, 0x41,
	0x7e, 0x5e, 0x71, 0xaa, 0x12, 0x83, 0x01, 0xa3, 0x93, 0x5e, 0x94, 0x4e, 0x59, 0x66, 0x49, 0x6a,
	0x71, 0xb1, 0x5e, 0x66, 0xbe, 0x3e, 0x84, 0xa5, 0x9f, 0x9e, 0xaf, 0x5f, 0x56, 0xa2, 0x0f, 0x76,
	0xa5, 0x3e, 0x9a, 0x4f, 0x92, 0xd8, 0xc0, 0xc2, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1c,
	0x01, 0x4d, 0x17, 0xfa, 0x00, 0x00, 0x00,
}
