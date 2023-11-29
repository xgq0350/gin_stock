package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type CustProductLocationDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustProductLocationDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustProductLocationDetailLogic {
	return &CustProductLocationDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CustProductLocationDetailLogic) CustProductLocationDetail(in *inventory.NormalDetailRequest) (*inventory.ProductlocationDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.ProductlocationDetailResponse{}, nil
}
