package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderproductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderproductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderproductDetailLogic {
	return &OrderproductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderproductDetailLogic) OrderproductDetail(in *inventory.NormalDetailRequest) (*inventory.OrderproductDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.OrderproductDetailResponse{}, nil
}
