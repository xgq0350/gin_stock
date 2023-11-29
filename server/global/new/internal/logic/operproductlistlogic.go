package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperproductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOperproductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperproductListLogic {
	return &OperproductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OperproductListLogic) OperproductList(in *inventory.NormalListRequest) (*inventory.OperproductListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.OperproductListResponse{}, nil
}
