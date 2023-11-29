package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type LocateproductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLocateproductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LocateproductListLogic {
	return &LocateproductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LocateproductListLogic) LocateproductList(in *inventory.NormalListRequest) (*inventory.LocateproductListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.LocateproductListResponse{}, nil
}
