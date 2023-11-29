package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type CustProductStoreDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustProductStoreDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustProductStoreDetailLogic {
	return &CustProductStoreDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CustProductStoreDetailLogic) CustProductStoreDetail(in *inventory.NormalDetailRequest) (*inventory.ProductlocationDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.ProductlocationDetailResponse{}, nil
}
