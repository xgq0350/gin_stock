package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderaddressUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderaddressUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderaddressUpdateLogic {
	return &OrderaddressUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderaddressUpdateLogic) OrderaddressUpdate(in *inventory.OrderaddressUpdateRequest) (*inventory.NormalResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.NormalResponse{}, nil
}
