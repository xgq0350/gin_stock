package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderaddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderaddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderaddressListLogic {
	return &OrderaddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderaddressListLogic) OrderaddressList(in *inventory.NormalListRequest) (*inventory.OrderaddressListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.OrderaddressListResponse{}, nil
}
