package service

import (
	"context"

	"mall-go/module/payment/service/internal/biz"
)

type PaymentService struct {
	UnimplementedPayServer
	ou *biz.PayUsecase
}

func (p PaymentService) Mp(ctx context.Context, request *MpRequest) (*MpReply, error) {
	// TODO implement me
	panic("implement me")
}

func (p PaymentService) Mini(ctx context.Context, request *MpRequest) (*MpReply, error) {
	// TODO implement me
	panic("implement me")
}

func (p PaymentService) App(ctx context.Context, request *AppRequest) (*AppReply, error) {
	// TODO implement me
	panic("implement me")
}

func (p PaymentService) Scan(ctx context.Context, request *ScanRequest) (*ScanReply, error) {
	// TODO implement me
	panic("implement me")
}

func (p PaymentService) Query(ctx context.Context, request *QueryRequest) (*QueryReply, error) {
	// TODO implement me
	panic("implement me")
}

func (p PaymentService) QueryRefund(ctx context.Context, request *QueryRefundRequest) (*QueryRefundReply, error) {
	// TODO implement me
	panic("implement me")
}

func (p PaymentService) Refund(ctx context.Context, request *RefundRequest) (*RefundReply, error) {
	// TODO implement me
	panic("implement me")
}

func (p PaymentService) Notify(ctx context.Context, request *NotifyRequest) (*NotifyReply, error) {
	// TODO implement me
	panic("implement me")
}

func (p PaymentService) NotifyRefund(ctx context.Context, request *NotifyRefundRequest) (*NotifyRefundReply, error) {
	// TODO implement me
	panic("implement me")
}

func (p PaymentService) mustEmbedUnimplementedPayServer() {
	// TODO implement me
	panic("implement me")
}

func NewPayService(ou *biz.PayUsecase) *PaymentService {
	return &PaymentService{
		ou: ou,
	}
}
