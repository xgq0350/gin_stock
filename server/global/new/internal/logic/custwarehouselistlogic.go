package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type CustwarehouseListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustwarehouseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustwarehouseListLogic {
	return &CustwarehouseListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CustwarehouseListLogic) CustwarehouseList(in *inventory.NormalListRequest) (*inventory.CustwarehouseListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.CustwarehouseListResponse{}, nil
}
