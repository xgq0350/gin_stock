package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderaddressDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderaddressDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderaddressDetailLogic {
	return &OrderaddressDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderaddressDetailLogic) OrderaddressDetail(in *inventory.NormalDetailRequest) (*inventory.OrderaddressDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.OrderaddressDetailResponse{}, nil
}
