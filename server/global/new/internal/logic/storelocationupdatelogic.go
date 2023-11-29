package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type StorelocationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStorelocationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StorelocationUpdateLogic {
	return &StorelocationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StorelocationUpdateLogic) StorelocationUpdate(in *inventory.StorelocationUpdateRequest) (*inventory.StorelocationUpdateResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.StorelocationUpdateResponse{}, nil
}
