// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

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

// WechatPayClient is the client API for WechatPay service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WechatPayClient interface {
	Mp(ctx context.Context, in *MpRequest, opts ...grpc.CallOption) (*MpReply, error)
	Mini(ctx context.Context, in *MpRequest, opts ...grpc.CallOption) (*MpReply, error)
	App(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppReply, error)
	Scan(ctx context.Context, in *ScanRequest, opts ...grpc.CallOption) (*ScanReply, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error)
	QueryRefund(ctx context.Context, in *QueryRefundRequest, opts ...grpc.CallOption) (*QueryRefundReply, error)
	Refund(ctx context.Context, in *RefundRequest, opts ...grpc.CallOption) (*RefundReply, error)
	Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyReply, error)
	NotifyRefund(ctx context.Context, in *NotifyRefundRequest, opts ...grpc.CallOption) (*NotifyRefundReply, error)
}

type wechatPayClient struct {
	cc grpc.ClientConnInterface
}

func NewWechatPayClient(cc grpc.ClientConnInterface) WechatPayClient {
	return &wechatPayClient{cc}
}

func (c *wechatPayClient) Mp(ctx context.Context, in *MpRequest, opts ...grpc.CallOption) (*MpReply, error) {
	out := new(MpReply)
	err := c.cc.Invoke(ctx, "/payment.service.WechatPay/Mp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatPayClient) Mini(ctx context.Context, in *MpRequest, opts ...grpc.CallOption) (*MpReply, error) {
	out := new(MpReply)
	err := c.cc.Invoke(ctx, "/payment.service.WechatPay/Mini", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatPayClient) App(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppReply, error) {
	out := new(AppReply)
	err := c.cc.Invoke(ctx, "/payment.service.WechatPay/App", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatPayClient) Scan(ctx context.Context, in *ScanRequest, opts ...grpc.CallOption) (*ScanReply, error) {
	out := new(ScanReply)
	err := c.cc.Invoke(ctx, "/payment.service.WechatPay/Scan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatPayClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error) {
	out := new(QueryReply)
	err := c.cc.Invoke(ctx, "/payment.service.WechatPay/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatPayClient) QueryRefund(ctx context.Context, in *QueryRefundRequest, opts ...grpc.CallOption) (*QueryRefundReply, error) {
	out := new(QueryRefundReply)
	err := c.cc.Invoke(ctx, "/payment.service.WechatPay/QueryRefund", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatPayClient) Refund(ctx context.Context, in *RefundRequest, opts ...grpc.CallOption) (*RefundReply, error) {
	out := new(RefundReply)
	err := c.cc.Invoke(ctx, "/payment.service.WechatPay/Refund", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatPayClient) Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyReply, error) {
	out := new(NotifyReply)
	err := c.cc.Invoke(ctx, "/payment.service.WechatPay/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatPayClient) NotifyRefund(ctx context.Context, in *NotifyRefundRequest, opts ...grpc.CallOption) (*NotifyRefundReply, error) {
	out := new(NotifyRefundReply)
	err := c.cc.Invoke(ctx, "/payment.service.WechatPay/NotifyRefund", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WechatPayServer is the server API for WechatPay service.
// All implementations must embed UnimplementedWechatPayServer
// for forward compatibility
type WechatPayServer interface {
	Mp(context.Context, *MpRequest) (*MpReply, error)
	Mini(context.Context, *MpRequest) (*MpReply, error)
	App(context.Context, *AppRequest) (*AppReply, error)
	Scan(context.Context, *ScanRequest) (*ScanReply, error)
	Query(context.Context, *QueryRequest) (*QueryReply, error)
	QueryRefund(context.Context, *QueryRefundRequest) (*QueryRefundReply, error)
	Refund(context.Context, *RefundRequest) (*RefundReply, error)
	Notify(context.Context, *NotifyRequest) (*NotifyReply, error)
	NotifyRefund(context.Context, *NotifyRefundRequest) (*NotifyRefundReply, error)
	mustEmbedUnimplementedWechatPayServer()
}

// UnimplementedWechatPayServer must be embedded to have forward compatible implementations.
type UnimplementedWechatPayServer struct {
}

func (UnimplementedWechatPayServer) Mp(context.Context, *MpRequest) (*MpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mp not implemented")
}
func (UnimplementedWechatPayServer) Mini(context.Context, *MpRequest) (*MpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mini not implemented")
}
func (UnimplementedWechatPayServer) App(context.Context, *AppRequest) (*AppReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method App not implemented")
}
func (UnimplementedWechatPayServer) Scan(context.Context, *ScanRequest) (*ScanReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Scan not implemented")
}
func (UnimplementedWechatPayServer) Query(context.Context, *QueryRequest) (*QueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedWechatPayServer) QueryRefund(context.Context, *QueryRefundRequest) (*QueryRefundReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRefund not implemented")
}
func (UnimplementedWechatPayServer) Refund(context.Context, *RefundRequest) (*RefundReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refund not implemented")
}
func (UnimplementedWechatPayServer) Notify(context.Context, *NotifyRequest) (*NotifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedWechatPayServer) NotifyRefund(context.Context, *NotifyRefundRequest) (*NotifyRefundReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyRefund not implemented")
}
func (UnimplementedWechatPayServer) mustEmbedUnimplementedWechatPayServer() {}

// UnsafeWechatPayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WechatPayServer will
// result in compilation errors.
type UnsafeWechatPayServer interface {
	mustEmbedUnimplementedWechatPayServer()
}

func RegisterWechatPayServer(s grpc.ServiceRegistrar, srv WechatPayServer) {
	s.RegisterService(&WechatPay_ServiceDesc, srv)
}

func _WechatPay_Mp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServer).Mp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.service.WechatPay/Mp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServer).Mp(ctx, req.(*MpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatPay_Mini_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServer).Mini(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.service.WechatPay/Mini",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServer).Mini(ctx, req.(*MpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatPay_App_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServer).App(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.service.WechatPay/App",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServer).App(ctx, req.(*AppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatPay_Scan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServer).Scan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.service.WechatPay/Scan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServer).Scan(ctx, req.(*ScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatPay_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.service.WechatPay/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatPay_QueryRefund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRefundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServer).QueryRefund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.service.WechatPay/QueryRefund",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServer).QueryRefund(ctx, req.(*QueryRefundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatPay_Refund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServer).Refund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.service.WechatPay/Refund",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServer).Refund(ctx, req.(*RefundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatPay_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.service.WechatPay/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServer).Notify(ctx, req.(*NotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatPay_NotifyRefund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRefundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServer).NotifyRefund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.service.WechatPay/NotifyRefund",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServer).NotifyRefund(ctx, req.(*NotifyRefundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WechatPay_ServiceDesc is the grpc.ServiceDesc for WechatPay service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WechatPay_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.service.WechatPay",
	HandlerType: (*WechatPayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Mp",
			Handler:    _WechatPay_Mp_Handler,
		},
		{
			MethodName: "Mini",
			Handler:    _WechatPay_Mini_Handler,
		},
		{
			MethodName: "App",
			Handler:    _WechatPay_App_Handler,
		},
		{
			MethodName: "Scan",
			Handler:    _WechatPay_Scan_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _WechatPay_Query_Handler,
		},
		{
			MethodName: "QueryRefund",
			Handler:    _WechatPay_QueryRefund_Handler,
		},
		{
			MethodName: "Refund",
			Handler:    _WechatPay_Refund_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _WechatPay_Notify_Handler,
		},
		{
			MethodName: "NotifyRefund",
			Handler:    _WechatPay_NotifyRefund_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/payment/service/wechatpay.proto",
}
