package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type CustwarehouseListExLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustwarehouseListExLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustwarehouseListExLogic {
	return &CustwarehouseListExLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CustwarehouseListExLogic) CustwarehouseListEx(in *inventory.NormalListRequest) (*inventory.CustwarehouseListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.CustwarehouseListResponse{}, nil
}
