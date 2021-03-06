// Code generated by protoc-gen-go. DO NOT EDIT.
// source: heat.proto

/*
Package heat is a generated protocol buffer package.

It is generated from these files:
	heat.proto

It has these top-level messages:
	Temp
*/
package heat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Temp struct {
	Value int32  `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Room  string `protobuf:"bytes,2,opt,name=room" json:"room,omitempty"`
}

func (m *Temp) Reset()                    { *m = Temp{} }
func (m *Temp) String() string            { return proto.CompactTextString(m) }
func (*Temp) ProtoMessage()               {}
func (*Temp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Temp) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Temp) GetRoom() string {
	if m != nil {
		return m.Room
	}
	return ""
}

func init() {
	proto.RegisterType((*Temp)(nil), "heat.Temp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Heat service

type HeatClient interface {
	Set(ctx context.Context, in *Temp, opts ...grpc.CallOption) (*Temp, error)
}

type heatClient struct {
	cc *grpc.ClientConn
}

func NewHeatClient(cc *grpc.ClientConn) HeatClient {
	return &heatClient{cc}
}

func (c *heatClient) Set(ctx context.Context, in *Temp, opts ...grpc.CallOption) (*Temp, error) {
	out := new(Temp)
	err := grpc.Invoke(ctx, "/heat.Heat/Set", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Heat service

type HeatServer interface {
	Set(context.Context, *Temp) (*Temp, error)
}

func RegisterHeatServer(s *grpc.Server, srv HeatServer) {
	s.RegisterService(&_Heat_serviceDesc, srv)
}

func _Heat_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Temp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeatServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heat.Heat/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeatServer).Set(ctx, req.(*Temp))
	}
	return interceptor(ctx, in, info, handler)
}

var _Heat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heat.Heat",
	HandlerType: (*HeatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _Heat_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heat.proto",
}

func init() { proto.RegisterFile("heat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x48, 0x4d, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x0c, 0xb8, 0x58, 0x42, 0x52,
	0x73, 0x0b, 0x84, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x58, 0x83, 0x20, 0x1c, 0x21, 0x21, 0x2e, 0x96, 0xa2, 0xfc, 0xfc, 0x5c, 0x09, 0x26, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x30, 0xdb, 0x48, 0x9d, 0x8b, 0xc5, 0x23, 0x35, 0xb1, 0x44, 0x48, 0x9e, 0x8b,
	0x39, 0x38, 0xb5, 0x44, 0x88, 0x4b, 0x0f, 0x6c, 0x26, 0xc8, 0x10, 0x29, 0x24, 0xb6, 0x12, 0x43,
	0x12, 0x1b, 0xd8, 0x1e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0x95, 0xd1, 0xa9, 0x75,
	0x00, 0x00, 0x00,
}
