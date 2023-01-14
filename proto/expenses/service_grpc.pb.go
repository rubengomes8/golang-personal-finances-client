// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/expenses/service.proto

package expenses

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

// ExpensesServiceClient is the client API for ExpensesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExpensesServiceClient interface {
	CreateExpense(ctx context.Context, in *ExpenseCreateRequest, opts ...grpc.CallOption) (*ExpenseCreateResponse, error)
	CreateExpenses(ctx context.Context, in *ExpensesCreateRequest, opts ...grpc.CallOption) (*ExpensesCreateResponse, error)
	GetExpensesByDate(ctx context.Context, in *ExpensesGetRequestByDate, opts ...grpc.CallOption) (*ExpensesGetResponse, error)
	GetExpensesByCategory(ctx context.Context, in *ExpensesGetRequestByCategory, opts ...grpc.CallOption) (*ExpensesGetResponse, error)
	GetExpensesBySubCategory(ctx context.Context, in *ExpensesGetRequestBySubCategory, opts ...grpc.CallOption) (*ExpensesGetResponse, error)
	GetExpensesByCard(ctx context.Context, in *ExpensesGetRequestByCard, opts ...grpc.CallOption) (*ExpensesGetResponse, error)
}

type expensesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExpensesServiceClient(cc grpc.ClientConnInterface) ExpensesServiceClient {
	return &expensesServiceClient{cc}
}

func (c *expensesServiceClient) CreateExpense(ctx context.Context, in *ExpenseCreateRequest, opts ...grpc.CallOption) (*ExpenseCreateResponse, error) {
	out := new(ExpenseCreateResponse)
	err := c.cc.Invoke(ctx, "/expenses.ExpensesService/CreateExpense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expensesServiceClient) CreateExpenses(ctx context.Context, in *ExpensesCreateRequest, opts ...grpc.CallOption) (*ExpensesCreateResponse, error) {
	out := new(ExpensesCreateResponse)
	err := c.cc.Invoke(ctx, "/expenses.ExpensesService/CreateExpenses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expensesServiceClient) GetExpensesByDate(ctx context.Context, in *ExpensesGetRequestByDate, opts ...grpc.CallOption) (*ExpensesGetResponse, error) {
	out := new(ExpensesGetResponse)
	err := c.cc.Invoke(ctx, "/expenses.ExpensesService/GetExpensesByDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expensesServiceClient) GetExpensesByCategory(ctx context.Context, in *ExpensesGetRequestByCategory, opts ...grpc.CallOption) (*ExpensesGetResponse, error) {
	out := new(ExpensesGetResponse)
	err := c.cc.Invoke(ctx, "/expenses.ExpensesService/GetExpensesByCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expensesServiceClient) GetExpensesBySubCategory(ctx context.Context, in *ExpensesGetRequestBySubCategory, opts ...grpc.CallOption) (*ExpensesGetResponse, error) {
	out := new(ExpensesGetResponse)
	err := c.cc.Invoke(ctx, "/expenses.ExpensesService/GetExpensesBySubCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expensesServiceClient) GetExpensesByCard(ctx context.Context, in *ExpensesGetRequestByCard, opts ...grpc.CallOption) (*ExpensesGetResponse, error) {
	out := new(ExpensesGetResponse)
	err := c.cc.Invoke(ctx, "/expenses.ExpensesService/GetExpensesByCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpensesServiceServer is the server API for ExpensesService service.
// All implementations must embed UnimplementedExpensesServiceServer
// for forward compatibility
type ExpensesServiceServer interface {
	CreateExpense(context.Context, *ExpenseCreateRequest) (*ExpenseCreateResponse, error)
	CreateExpenses(context.Context, *ExpensesCreateRequest) (*ExpensesCreateResponse, error)
	GetExpensesByDate(context.Context, *ExpensesGetRequestByDate) (*ExpensesGetResponse, error)
	GetExpensesByCategory(context.Context, *ExpensesGetRequestByCategory) (*ExpensesGetResponse, error)
	GetExpensesBySubCategory(context.Context, *ExpensesGetRequestBySubCategory) (*ExpensesGetResponse, error)
	GetExpensesByCard(context.Context, *ExpensesGetRequestByCard) (*ExpensesGetResponse, error)
	mustEmbedUnimplementedExpensesServiceServer()
}

// UnimplementedExpensesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExpensesServiceServer struct {
}

func (UnimplementedExpensesServiceServer) CreateExpense(context.Context, *ExpenseCreateRequest) (*ExpenseCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExpense not implemented")
}
func (UnimplementedExpensesServiceServer) CreateExpenses(context.Context, *ExpensesCreateRequest) (*ExpensesCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExpenses not implemented")
}
func (UnimplementedExpensesServiceServer) GetExpensesByDate(context.Context, *ExpensesGetRequestByDate) (*ExpensesGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExpensesByDate not implemented")
}
func (UnimplementedExpensesServiceServer) GetExpensesByCategory(context.Context, *ExpensesGetRequestByCategory) (*ExpensesGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExpensesByCategory not implemented")
}
func (UnimplementedExpensesServiceServer) GetExpensesBySubCategory(context.Context, *ExpensesGetRequestBySubCategory) (*ExpensesGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExpensesBySubCategory not implemented")
}
func (UnimplementedExpensesServiceServer) GetExpensesByCard(context.Context, *ExpensesGetRequestByCard) (*ExpensesGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExpensesByCard not implemented")
}
func (UnimplementedExpensesServiceServer) mustEmbedUnimplementedExpensesServiceServer() {}

// UnsafeExpensesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExpensesServiceServer will
// result in compilation errors.
type UnsafeExpensesServiceServer interface {
	mustEmbedUnimplementedExpensesServiceServer()
}

func RegisterExpensesServiceServer(s grpc.ServiceRegistrar, srv ExpensesServiceServer) {
	s.RegisterService(&ExpensesService_ServiceDesc, srv)
}

func _ExpensesService_CreateExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpenseCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpensesServiceServer).CreateExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/expenses.ExpensesService/CreateExpense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpensesServiceServer).CreateExpense(ctx, req.(*ExpenseCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpensesService_CreateExpenses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpensesCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpensesServiceServer).CreateExpenses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/expenses.ExpensesService/CreateExpenses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpensesServiceServer).CreateExpenses(ctx, req.(*ExpensesCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpensesService_GetExpensesByDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpensesGetRequestByDate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpensesServiceServer).GetExpensesByDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/expenses.ExpensesService/GetExpensesByDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpensesServiceServer).GetExpensesByDate(ctx, req.(*ExpensesGetRequestByDate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpensesService_GetExpensesByCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpensesGetRequestByCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpensesServiceServer).GetExpensesByCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/expenses.ExpensesService/GetExpensesByCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpensesServiceServer).GetExpensesByCategory(ctx, req.(*ExpensesGetRequestByCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpensesService_GetExpensesBySubCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpensesGetRequestBySubCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpensesServiceServer).GetExpensesBySubCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/expenses.ExpensesService/GetExpensesBySubCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpensesServiceServer).GetExpensesBySubCategory(ctx, req.(*ExpensesGetRequestBySubCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpensesService_GetExpensesByCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpensesGetRequestByCard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpensesServiceServer).GetExpensesByCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/expenses.ExpensesService/GetExpensesByCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpensesServiceServer).GetExpensesByCard(ctx, req.(*ExpensesGetRequestByCard))
	}
	return interceptor(ctx, in, info, handler)
}

// ExpensesService_ServiceDesc is the grpc.ServiceDesc for ExpensesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExpensesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "expenses.ExpensesService",
	HandlerType: (*ExpensesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExpense",
			Handler:    _ExpensesService_CreateExpense_Handler,
		},
		{
			MethodName: "CreateExpenses",
			Handler:    _ExpensesService_CreateExpenses_Handler,
		},
		{
			MethodName: "GetExpensesByDate",
			Handler:    _ExpensesService_GetExpensesByDate_Handler,
		},
		{
			MethodName: "GetExpensesByCategory",
			Handler:    _ExpensesService_GetExpensesByCategory_Handler,
		},
		{
			MethodName: "GetExpensesBySubCategory",
			Handler:    _ExpensesService_GetExpensesBySubCategory_Handler,
		},
		{
			MethodName: "GetExpensesByCard",
			Handler:    _ExpensesService_GetExpensesByCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/expenses/service.proto",
}
