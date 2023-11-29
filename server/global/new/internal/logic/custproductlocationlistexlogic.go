package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type CustProductLocationListExLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustProductLocationListExLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustProductLocationListExLogic {
	return &CustProductLocationListExLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CustProductLocationListExLogic) CustProductLocationListEx(in *inventory.NormalListRequest) (*inventory.ProductlocationListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.ProductlocationListResponse{}, nil
}
