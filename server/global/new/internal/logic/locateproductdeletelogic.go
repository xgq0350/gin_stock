package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type LocateproductDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLocateproductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LocateproductDeleteLogic {
	return &LocateproductDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LocateproductDeleteLogic) LocateproductDelete(in *inventory.NormalDeleteRequest) (*inventory.NormalResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.NormalResponse{}, nil
}
