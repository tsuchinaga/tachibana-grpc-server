// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: tachibanapb/tachibana.proto

package tachibanapb

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

// TachibanaServiceClient is the client API for TachibanaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TachibanaServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	NewOrder(ctx context.Context, in *NewOrderRequest, opts ...grpc.CallOption) (*NewOrderResponse, error)
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
	OrderList(ctx context.Context, in *OrderListRequest, opts ...grpc.CallOption) (*OrderListResponse, error)
	OrderDetail(ctx context.Context, in *OrderDetailRequest, opts ...grpc.CallOption) (*OrderDetailResponse, error)
	StockMaster(ctx context.Context, in *StockMasterRequest, opts ...grpc.CallOption) (*StockMasterResponse, error)
	StockExchangeMaster(ctx context.Context, in *StockExchangeMasterRequest, opts ...grpc.CallOption) (*StockExchangeMasterResponse, error)
	MarketPrice(ctx context.Context, in *MarketPriceRequest, opts ...grpc.CallOption) (*MarketPriceResponse, error)
	BusinessDay(ctx context.Context, in *BusinessDayRequest, opts ...grpc.CallOption) (*BusinessDayResponse, error)
	TickGroup(ctx context.Context, in *TickGroupRequest, opts ...grpc.CallOption) (*TickGroupResponse, error)
}

type tachibanaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTachibanaServiceClient(cc grpc.ClientConnInterface) TachibanaServiceClient {
	return &tachibanaServiceClient{cc}
}

func (c *tachibanaServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/tachibanapb.TachibanaService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tachibanaServiceClient) NewOrder(ctx context.Context, in *NewOrderRequest, opts ...grpc.CallOption) (*NewOrderResponse, error) {
	out := new(NewOrderResponse)
	err := c.cc.Invoke(ctx, "/tachibanapb.TachibanaService/NewOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tachibanaServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error) {
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, "/tachibanapb.TachibanaService/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tachibanaServiceClient) OrderList(ctx context.Context, in *OrderListRequest, opts ...grpc.CallOption) (*OrderListResponse, error) {
	out := new(OrderListResponse)
	err := c.cc.Invoke(ctx, "/tachibanapb.TachibanaService/OrderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tachibanaServiceClient) OrderDetail(ctx context.Context, in *OrderDetailRequest, opts ...grpc.CallOption) (*OrderDetailResponse, error) {
	out := new(OrderDetailResponse)
	err := c.cc.Invoke(ctx, "/tachibanapb.TachibanaService/OrderDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tachibanaServiceClient) StockMaster(ctx context.Context, in *StockMasterRequest, opts ...grpc.CallOption) (*StockMasterResponse, error) {
	out := new(StockMasterResponse)
	err := c.cc.Invoke(ctx, "/tachibanapb.TachibanaService/StockMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tachibanaServiceClient) StockExchangeMaster(ctx context.Context, in *StockExchangeMasterRequest, opts ...grpc.CallOption) (*StockExchangeMasterResponse, error) {
	out := new(StockExchangeMasterResponse)
	err := c.cc.Invoke(ctx, "/tachibanapb.TachibanaService/StockExchangeMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tachibanaServiceClient) MarketPrice(ctx context.Context, in *MarketPriceRequest, opts ...grpc.CallOption) (*MarketPriceResponse, error) {
	out := new(MarketPriceResponse)
	err := c.cc.Invoke(ctx, "/tachibanapb.TachibanaService/MarketPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tachibanaServiceClient) BusinessDay(ctx context.Context, in *BusinessDayRequest, opts ...grpc.CallOption) (*BusinessDayResponse, error) {
	out := new(BusinessDayResponse)
	err := c.cc.Invoke(ctx, "/tachibanapb.TachibanaService/BusinessDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tachibanaServiceClient) TickGroup(ctx context.Context, in *TickGroupRequest, opts ...grpc.CallOption) (*TickGroupResponse, error) {
	out := new(TickGroupResponse)
	err := c.cc.Invoke(ctx, "/tachibanapb.TachibanaService/TickGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TachibanaServiceServer is the server API for TachibanaService service.
// All implementations must embed UnimplementedTachibanaServiceServer
// for forward compatibility
type TachibanaServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	NewOrder(context.Context, *NewOrderRequest) (*NewOrderResponse, error)
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
	OrderList(context.Context, *OrderListRequest) (*OrderListResponse, error)
	OrderDetail(context.Context, *OrderDetailRequest) (*OrderDetailResponse, error)
	StockMaster(context.Context, *StockMasterRequest) (*StockMasterResponse, error)
	StockExchangeMaster(context.Context, *StockExchangeMasterRequest) (*StockExchangeMasterResponse, error)
	MarketPrice(context.Context, *MarketPriceRequest) (*MarketPriceResponse, error)
	BusinessDay(context.Context, *BusinessDayRequest) (*BusinessDayResponse, error)
	TickGroup(context.Context, *TickGroupRequest) (*TickGroupResponse, error)
	mustEmbedUnimplementedTachibanaServiceServer()
}

// UnimplementedTachibanaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTachibanaServiceServer struct {
}

func (UnimplementedTachibanaServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedTachibanaServiceServer) NewOrder(context.Context, *NewOrderRequest) (*NewOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewOrder not implemented")
}
func (UnimplementedTachibanaServiceServer) CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedTachibanaServiceServer) OrderList(context.Context, *OrderListRequest) (*OrderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderList not implemented")
}
func (UnimplementedTachibanaServiceServer) OrderDetail(context.Context, *OrderDetailRequest) (*OrderDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDetail not implemented")
}
func (UnimplementedTachibanaServiceServer) StockMaster(context.Context, *StockMasterRequest) (*StockMasterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StockMaster not implemented")
}
func (UnimplementedTachibanaServiceServer) StockExchangeMaster(context.Context, *StockExchangeMasterRequest) (*StockExchangeMasterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StockExchangeMaster not implemented")
}
func (UnimplementedTachibanaServiceServer) MarketPrice(context.Context, *MarketPriceRequest) (*MarketPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketPrice not implemented")
}
func (UnimplementedTachibanaServiceServer) BusinessDay(context.Context, *BusinessDayRequest) (*BusinessDayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BusinessDay not implemented")
}
func (UnimplementedTachibanaServiceServer) TickGroup(context.Context, *TickGroupRequest) (*TickGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TickGroup not implemented")
}
func (UnimplementedTachibanaServiceServer) mustEmbedUnimplementedTachibanaServiceServer() {}

// UnsafeTachibanaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TachibanaServiceServer will
// result in compilation errors.
type UnsafeTachibanaServiceServer interface {
	mustEmbedUnimplementedTachibanaServiceServer()
}

func RegisterTachibanaServiceServer(s grpc.ServiceRegistrar, srv TachibanaServiceServer) {
	s.RegisterService(&TachibanaService_ServiceDesc, srv)
}

func _TachibanaService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TachibanaServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tachibanapb.TachibanaService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TachibanaServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TachibanaService_NewOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TachibanaServiceServer).NewOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tachibanapb.TachibanaService/NewOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TachibanaServiceServer).NewOrder(ctx, req.(*NewOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TachibanaService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TachibanaServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tachibanapb.TachibanaService/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TachibanaServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TachibanaService_OrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TachibanaServiceServer).OrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tachibanapb.TachibanaService/OrderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TachibanaServiceServer).OrderList(ctx, req.(*OrderListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TachibanaService_OrderDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TachibanaServiceServer).OrderDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tachibanapb.TachibanaService/OrderDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TachibanaServiceServer).OrderDetail(ctx, req.(*OrderDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TachibanaService_StockMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockMasterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TachibanaServiceServer).StockMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tachibanapb.TachibanaService/StockMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TachibanaServiceServer).StockMaster(ctx, req.(*StockMasterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TachibanaService_StockExchangeMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockExchangeMasterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TachibanaServiceServer).StockExchangeMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tachibanapb.TachibanaService/StockExchangeMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TachibanaServiceServer).StockExchangeMaster(ctx, req.(*StockExchangeMasterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TachibanaService_MarketPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TachibanaServiceServer).MarketPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tachibanapb.TachibanaService/MarketPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TachibanaServiceServer).MarketPrice(ctx, req.(*MarketPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TachibanaService_BusinessDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BusinessDayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TachibanaServiceServer).BusinessDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tachibanapb.TachibanaService/BusinessDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TachibanaServiceServer).BusinessDay(ctx, req.(*BusinessDayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TachibanaService_TickGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TickGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TachibanaServiceServer).TickGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tachibanapb.TachibanaService/TickGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TachibanaServiceServer).TickGroup(ctx, req.(*TickGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TachibanaService_ServiceDesc is the grpc.ServiceDesc for TachibanaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TachibanaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tachibanapb.TachibanaService",
	HandlerType: (*TachibanaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _TachibanaService_Login_Handler,
		},
		{
			MethodName: "NewOrder",
			Handler:    _TachibanaService_NewOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _TachibanaService_CancelOrder_Handler,
		},
		{
			MethodName: "OrderList",
			Handler:    _TachibanaService_OrderList_Handler,
		},
		{
			MethodName: "OrderDetail",
			Handler:    _TachibanaService_OrderDetail_Handler,
		},
		{
			MethodName: "StockMaster",
			Handler:    _TachibanaService_StockMaster_Handler,
		},
		{
			MethodName: "StockExchangeMaster",
			Handler:    _TachibanaService_StockExchangeMaster_Handler,
		},
		{
			MethodName: "MarketPrice",
			Handler:    _TachibanaService_MarketPrice_Handler,
		},
		{
			MethodName: "BusinessDay",
			Handler:    _TachibanaService_BusinessDay_Handler,
		},
		{
			MethodName: "TickGroup",
			Handler:    _TachibanaService_TickGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tachibanapb/tachibana.proto",
}
