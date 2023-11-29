package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderproductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderproductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderproductListLogic {
	return &OrderproductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderproductListLogic) OrderproductList(in *inventory.NormalStatusListRequest) (*inventory.OrderproductListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.OrderproductListResponse{}, nil
}
