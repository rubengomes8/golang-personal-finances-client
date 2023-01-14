// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/cards/service.proto

package cards

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CardServiceClient is the client API for CardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardServiceClient interface {
	CreateCard(ctx context.Context, in *CardCreateRequest, opts ...grpc.CallOption) (*CardCreateResponse, error)
	GetCardByName(ctx context.Context, in *CardGetRequestByName, opts ...grpc.CallOption) (*CardGetResponse, error)
}

type cardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCardServiceClient(cc grpc.ClientConnInterface) CardServiceClient {
	return &cardServiceClient{cc}
}

func (c *cardServiceClient) CreateCard(ctx context.Context, in *CardCreateRequest, opts ...grpc.CallOption) (*CardCreateResponse, error) {
	out := new(CardCreateResponse)
	err := c.cc.Invoke(ctx, "/cards.CardService/CreateCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceClient) GetCardByName(ctx context.Context, in *CardGetRequestByName, opts ...grpc.CallOption) (*CardGetResponse, error) {
	out := new(CardGetResponse)
	err := c.cc.Invoke(ctx, "/cards.CardService/GetCardByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardServiceServer is the server API for CardService service.
// All implementations must embed UnimplementedCardServiceServer
// for forward compatibility
type CardServiceServer interface {
	CreateCard(context.Context, *CardCreateRequest) (*CardCreateResponse, error)
	GetCardByName(context.Context, *CardGetRequestByName) (*CardGetResponse, error)
	mustEmbedUnimplementedCardServiceServer()
}

// UnimplementedCardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCardServiceServer struct {
}

func (UnimplementedCardServiceServer) CreateCard(context.Context, *CardCreateRequest) (*CardCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCard not implemented")
}
func (UnimplementedCardServiceServer) GetCardByName(context.Context, *CardGetRequestByName) (*CardGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCardByName not implemented")
}
func (UnimplementedCardServiceServer) mustEmbedUnimplementedCardServiceServer() {}

// UnsafeCardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardServiceServer will
// result in compilation errors.
type UnsafeCardServiceServer interface {
	mustEmbedUnimplementedCardServiceServer()
}

func RegisterCardServiceServer(s grpc.ServiceRegistrar, srv CardServiceServer) {
	s.RegisterService(&CardService_ServiceDesc, srv)
}

func _CardService_CreateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).CreateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cards.CardService/CreateCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).CreateCard(ctx, req.(*CardCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardService_GetCardByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardGetRequestByName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).GetCardByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cards.CardService/GetCardByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).GetCardByName(ctx, req.(*CardGetRequestByName))
	}
	return interceptor(ctx, in, info, handler)
}

// CardService_ServiceDesc is the grpc.ServiceDesc for CardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cards.CardService",
	HandlerType: (*CardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCard",
			Handler:    _CardService_CreateCard_Handler,
		},
		{
			MethodName: "GetCardByName",
			Handler:    _CardService_GetCardByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cards/service.proto",
}
