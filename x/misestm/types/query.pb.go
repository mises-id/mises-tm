// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: misestm/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("misestm/query.proto", fileDescriptor_b032d76fc324b94e) }

var fileDescriptor_b032d76fc324b94e = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0xcd, 0x2c, 0x4e,
	0x2d, 0x2e, 0xc9, 0xd5, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x07, 0x0b, 0x66, 0xa6, 0xe8, 0x41, 0x25, 0x61, 0xb4, 0x94, 0x4c, 0x7a, 0x7e, 0x7e, 0x7a,
	0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49, 0x62, 0x49, 0x66, 0x7e,
	0x5e, 0x31, 0x44, 0x9b, 0x94, 0x56, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0xb1, 0x7e, 0x52, 0x62, 0x71,
	0x2a, 0xc4, 0x3c, 0xfd, 0x32, 0xc3, 0xa4, 0xd4, 0x92, 0x44, 0x43, 0xfd, 0x82, 0xc4, 0xf4, 0xcc,
	0x3c, 0xb0, 0x62, 0x88, 0x5a, 0x23, 0x76, 0x2e, 0xd6, 0x40, 0x90, 0x0a, 0x27, 0xb7, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b,
	0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x49, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0x07, 0x3b, 0x40, 0x37, 0x33, 0x05, 0xca, 0x28, 0xc9, 0xd5, 0xaf, 0xd0,
	0x87, 0xb9, 0xbc, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x6c, 0xae, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x51, 0x89, 0x8e, 0x12, 0xd1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "misesid.misestm.misestm.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "misestm/query.proto",
}
