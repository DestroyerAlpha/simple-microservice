// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: api/customer/service.proto

package customer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CustomerService_GetMenu_FullMethodName        = "/customer.CustomerService/GetMenu"
	CustomerService_PlaceFoodOrder_FullMethodName = "/customer.CustomerService/PlaceFoodOrder"
	CustomerService_ReviewFoodItem_FullMethodName = "/customer.CustomerService/ReviewFoodItem"
)

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerServiceClient interface {
	GetMenu(ctx context.Context, in *GetMenuRequest, opts ...grpc.CallOption) (*GetMenuResponse, error)
	PlaceFoodOrder(ctx context.Context, in *PlaceFoodOrderRequest, opts ...grpc.CallOption) (*PlaceFoodOrderResponse, error)
	ReviewFoodItem(ctx context.Context, in *ReviewFoodItemRequest, opts ...grpc.CallOption) (*ReviewFoodItemResponse, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) GetMenu(ctx context.Context, in *GetMenuRequest, opts ...grpc.CallOption) (*GetMenuResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMenuResponse)
	err := c.cc.Invoke(ctx, CustomerService_GetMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) PlaceFoodOrder(ctx context.Context, in *PlaceFoodOrderRequest, opts ...grpc.CallOption) (*PlaceFoodOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlaceFoodOrderResponse)
	err := c.cc.Invoke(ctx, CustomerService_PlaceFoodOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) ReviewFoodItem(ctx context.Context, in *ReviewFoodItemRequest, opts ...grpc.CallOption) (*ReviewFoodItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReviewFoodItemResponse)
	err := c.cc.Invoke(ctx, CustomerService_ReviewFoodItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
// All implementations must embed UnimplementedCustomerServiceServer
// for forward compatibility.
type CustomerServiceServer interface {
	GetMenu(context.Context, *GetMenuRequest) (*GetMenuResponse, error)
	PlaceFoodOrder(context.Context, *PlaceFoodOrderRequest) (*PlaceFoodOrderResponse, error)
	ReviewFoodItem(context.Context, *ReviewFoodItemRequest) (*ReviewFoodItemResponse, error)
	mustEmbedUnimplementedCustomerServiceServer()
}

// UnimplementedCustomerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCustomerServiceServer struct{}

func (UnimplementedCustomerServiceServer) GetMenu(context.Context, *GetMenuRequest) (*GetMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenu not implemented")
}
func (UnimplementedCustomerServiceServer) PlaceFoodOrder(context.Context, *PlaceFoodOrderRequest) (*PlaceFoodOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceFoodOrder not implemented")
}
func (UnimplementedCustomerServiceServer) ReviewFoodItem(context.Context, *ReviewFoodItemRequest) (*ReviewFoodItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReviewFoodItem not implemented")
}
func (UnimplementedCustomerServiceServer) mustEmbedUnimplementedCustomerServiceServer() {}
func (UnimplementedCustomerServiceServer) testEmbeddedByValue()                         {}

// UnsafeCustomerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServiceServer will
// result in compilation errors.
type UnsafeCustomerServiceServer interface {
	mustEmbedUnimplementedCustomerServiceServer()
}

func RegisterCustomerServiceServer(s grpc.ServiceRegistrar, srv CustomerServiceServer) {
	// If the following call pancis, it indicates UnimplementedCustomerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CustomerService_ServiceDesc, srv)
}

func _CustomerService_GetMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_GetMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetMenu(ctx, req.(*GetMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_PlaceFoodOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceFoodOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).PlaceFoodOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_PlaceFoodOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).PlaceFoodOrder(ctx, req.(*PlaceFoodOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_ReviewFoodItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewFoodItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).ReviewFoodItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_ReviewFoodItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).ReviewFoodItem(ctx, req.(*ReviewFoodItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerService_ServiceDesc is the grpc.ServiceDesc for CustomerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "customer.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMenu",
			Handler:    _CustomerService_GetMenu_Handler,
		},
		{
			MethodName: "PlaceFoodOrder",
			Handler:    _CustomerService_PlaceFoodOrder_Handler,
		},
		{
			MethodName: "ReviewFoodItem",
			Handler:    _CustomerService_ReviewFoodItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/customer/service.proto",
}
