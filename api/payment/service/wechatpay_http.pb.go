// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.5

package service

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type WechatPayHTTPServer interface {
	App(context.Context, *AppRequest) (*AppReply, error)
	Mini(context.Context, *MpRequest) (*MpReply, error)
	Mp(context.Context, *MpRequest) (*MpReply, error)
	Notify(context.Context, *NotifyRequest) (*NotifyReply, error)
	NotifyRefund(context.Context, *NotifyRefundRequest) (*NotifyRefundReply, error)
	Query(context.Context, *QueryRequest) (*QueryReply, error)
	QueryRefund(context.Context, *QueryRefundRequest) (*QueryRefundReply, error)
	Refund(context.Context, *RefundRequest) (*RefundReply, error)
	Scan(context.Context, *ScanRequest) (*ScanReply, error)
}

func RegisterWechatPayHTTPServer(s *http.Server, srv WechatPayHTTPServer) {
	r := s.Route("/")
	r.POST("/mp", _WechatPay_Mp0_HTTP_Handler(srv))
	r.POST("/mini", _WechatPay_Mini0_HTTP_Handler(srv))
	r.POST("/app", _WechatPay_App0_HTTP_Handler(srv))
	r.POST("/scan", _WechatPay_Scan0_HTTP_Handler(srv))
	r.POST("/query", _WechatPay_Query0_HTTP_Handler(srv))
	r.POST("/query/refund", _WechatPay_QueryRefund0_HTTP_Handler(srv))
	r.POST("/refund", _WechatPay_Refund0_HTTP_Handler(srv))
	r.POST("/notify", _WechatPay_Notify0_HTTP_Handler(srv))
	r.POST("/notify/refund", _WechatPay_NotifyRefund0_HTTP_Handler(srv))
}

func _WechatPay_Mp0_HTTP_Handler(srv WechatPayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MpRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/payment.service.WechatPay/Mp")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Mp(ctx, req.(*MpRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MpReply)
		return ctx.Result(200, reply)
	}
}

func _WechatPay_Mini0_HTTP_Handler(srv WechatPayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MpRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/payment.service.WechatPay/Mini")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Mini(ctx, req.(*MpRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MpReply)
		return ctx.Result(200, reply)
	}
}

func _WechatPay_App0_HTTP_Handler(srv WechatPayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/payment.service.WechatPay/App")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.App(ctx, req.(*AppRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AppReply)
		return ctx.Result(200, reply)
	}
}

func _WechatPay_Scan0_HTTP_Handler(srv WechatPayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ScanRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/payment.service.WechatPay/Scan")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Scan(ctx, req.(*ScanRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ScanReply)
		return ctx.Result(200, reply)
	}
}

func _WechatPay_Query0_HTTP_Handler(srv WechatPayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in QueryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/payment.service.WechatPay/Query")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Query(ctx, req.(*QueryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*QueryReply)
		return ctx.Result(200, reply)
	}
}

func _WechatPay_QueryRefund0_HTTP_Handler(srv WechatPayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in QueryRefundRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/payment.service.WechatPay/QueryRefund")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.QueryRefund(ctx, req.(*QueryRefundRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*QueryRefundReply)
		return ctx.Result(200, reply)
	}
}

func _WechatPay_Refund0_HTTP_Handler(srv WechatPayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RefundRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/payment.service.WechatPay/Refund")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Refund(ctx, req.(*RefundRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RefundReply)
		return ctx.Result(200, reply)
	}
}

func _WechatPay_Notify0_HTTP_Handler(srv WechatPayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NotifyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/payment.service.WechatPay/Notify")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Notify(ctx, req.(*NotifyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NotifyReply)
		return ctx.Result(200, reply)
	}
}

func _WechatPay_NotifyRefund0_HTTP_Handler(srv WechatPayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NotifyRefundRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/payment.service.WechatPay/NotifyRefund")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.NotifyRefund(ctx, req.(*NotifyRefundRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NotifyRefundReply)
		return ctx.Result(200, reply)
	}
}

type WechatPayHTTPClient interface {
	App(ctx context.Context, req *AppRequest, opts ...http.CallOption) (rsp *AppReply, err error)
	Mini(ctx context.Context, req *MpRequest, opts ...http.CallOption) (rsp *MpReply, err error)
	Mp(ctx context.Context, req *MpRequest, opts ...http.CallOption) (rsp *MpReply, err error)
	Notify(ctx context.Context, req *NotifyRequest, opts ...http.CallOption) (rsp *NotifyReply, err error)
	NotifyRefund(ctx context.Context, req *NotifyRefundRequest, opts ...http.CallOption) (rsp *NotifyRefundReply, err error)
	Query(ctx context.Context, req *QueryRequest, opts ...http.CallOption) (rsp *QueryReply, err error)
	QueryRefund(ctx context.Context, req *QueryRefundRequest, opts ...http.CallOption) (rsp *QueryRefundReply, err error)
	Refund(ctx context.Context, req *RefundRequest, opts ...http.CallOption) (rsp *RefundReply, err error)
	Scan(ctx context.Context, req *ScanRequest, opts ...http.CallOption) (rsp *ScanReply, err error)
}

type WechatPayHTTPClientImpl struct {
	cc *http.Client
}

func NewWechatPayHTTPClient(client *http.Client) WechatPayHTTPClient {
	return &WechatPayHTTPClientImpl{client}
}

func (c *WechatPayHTTPClientImpl) App(ctx context.Context, in *AppRequest, opts ...http.CallOption) (*AppReply, error) {
	var out AppReply
	pattern := "/app"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/payment.service.WechatPay/App"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WechatPayHTTPClientImpl) Mini(ctx context.Context, in *MpRequest, opts ...http.CallOption) (*MpReply, error) {
	var out MpReply
	pattern := "/mini"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/payment.service.WechatPay/Mini"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WechatPayHTTPClientImpl) Mp(ctx context.Context, in *MpRequest, opts ...http.CallOption) (*MpReply, error) {
	var out MpReply
	pattern := "/mp"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/payment.service.WechatPay/Mp"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WechatPayHTTPClientImpl) Notify(ctx context.Context, in *NotifyRequest, opts ...http.CallOption) (*NotifyReply, error) {
	var out NotifyReply
	pattern := "/notify"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/payment.service.WechatPay/Notify"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WechatPayHTTPClientImpl) NotifyRefund(ctx context.Context, in *NotifyRefundRequest, opts ...http.CallOption) (*NotifyRefundReply, error) {
	var out NotifyRefundReply
	pattern := "/notify/refund"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/payment.service.WechatPay/NotifyRefund"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WechatPayHTTPClientImpl) Query(ctx context.Context, in *QueryRequest, opts ...http.CallOption) (*QueryReply, error) {
	var out QueryReply
	pattern := "/query"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/payment.service.WechatPay/Query"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WechatPayHTTPClientImpl) QueryRefund(ctx context.Context, in *QueryRefundRequest, opts ...http.CallOption) (*QueryRefundReply, error) {
	var out QueryRefundReply
	pattern := "/query/refund"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/payment.service.WechatPay/QueryRefund"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WechatPayHTTPClientImpl) Refund(ctx context.Context, in *RefundRequest, opts ...http.CallOption) (*RefundReply, error) {
	var out RefundReply
	pattern := "/refund"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/payment.service.WechatPay/Refund"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WechatPayHTTPClientImpl) Scan(ctx context.Context, in *ScanRequest, opts ...http.CallOption) (*ScanReply, error) {
	var out ScanReply
	pattern := "/scan"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/payment.service.WechatPay/Scan"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
