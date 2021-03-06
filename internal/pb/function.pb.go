// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/protobuf-spec/function.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("api/protobuf-spec/function.proto", fileDescriptor_949a68f1f25eb768) }

var fileDescriptor_949a68f1f25eb768 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xcd, 0x31, 0x0b, 0xc2, 0x30,
	0x10, 0xc5, 0x71, 0xa1, 0x20, 0xd2, 0x49, 0xea, 0xd6, 0x49, 0xdc, 0x9b, 0x80, 0x22, 0x82, 0x9b,
	0x0a, 0xba, 0x4a, 0x47, 0xb7, 0x24, 0x5e, 0xd3, 0x40, 0x92, 0x0b, 0xc9, 0xdd, 0xf7, 0x17, 0xaa,
	0xe8, 0xe2, 0xfa, 0xe7, 0xf7, 0x78, 0xf5, 0x5a, 0x25, 0x27, 0x53, 0x46, 0x42, 0xcd, 0x43, 0x57,
	0x12, 0x18, 0x39, 0x70, 0x34, 0xe4, 0x30, 0x8a, 0x29, 0x37, 0x95, 0x4a, 0xae, 0xfd, 0xc3, 0x02,
	0x94, 0xa2, 0x2c, 0x94, 0x37, 0xdb, 0x1e, 0xeb, 0xc5, 0xf5, 0x33, 0x6c, 0x44, 0x5d, 0xf5, 0x1c,
	0x9b, 0x95, 0xf8, 0x9a, 0x53, 0xb6, 0x1c, 0x20, 0x52, 0x69, 0x97, 0xbf, 0xd8, 0x43, 0x61, 0x4f,
	0x9b, 0xd9, 0xf9, 0xf0, 0xd8, 0x5b, 0x47, 0x23, 0x6b, 0x61, 0x30, 0xc8, 0x1b, 0xa2, 0xf5, 0x70,
	0xf1, 0xc8, 0xcf, 0xbb, 0x57, 0x34, 0x60, 0x0e, 0x12, 0x13, 0xc4, 0x2e, 0x28, 0x32, 0xa3, 0x74,
	0x91, 0x20, 0x47, 0xe5, 0x65, 0xd2, 0x7a, 0x3e, 0x7d, 0xef, 0x5e, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x72, 0xee, 0x1f, 0xa8, 0xc6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FunctionClient is the client API for Function service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FunctionClient interface {
	// The assumption is that there will be one service for each MMF that is
	// being served.  Build your MMF in the appropriate serving harness, deploy it
	// to the K8s cluster with a unique service name, then connect to that service
	// and call 'Run()' to execute the fuction.
	Run(ctx context.Context, in *Arguments, opts ...grpc.CallOption) (*Result, error)
}

type functionClient struct {
	cc *grpc.ClientConn
}

func NewFunctionClient(cc *grpc.ClientConn) FunctionClient {
	return &functionClient{cc}
}

func (c *functionClient) Run(ctx context.Context, in *Arguments, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/api.Function/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FunctionServer is the server API for Function service.
type FunctionServer interface {
	// The assumption is that there will be one service for each MMF that is
	// being served.  Build your MMF in the appropriate serving harness, deploy it
	// to the K8s cluster with a unique service name, then connect to that service
	// and call 'Run()' to execute the fuction.
	Run(context.Context, *Arguments) (*Result, error)
}

func RegisterFunctionServer(s *grpc.Server, srv FunctionServer) {
	s.RegisterService(&_Function_serviceDesc, srv)
}

func _Function_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Arguments)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Function/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionServer).Run(ctx, req.(*Arguments))
	}
	return interceptor(ctx, in, info, handler)
}

var _Function_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Function",
	HandlerType: (*FunctionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _Function_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/protobuf-spec/function.proto",
}
