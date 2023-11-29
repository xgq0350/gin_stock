package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type StorelocationDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStorelocationDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StorelocationDetailLogic {
	return &StorelocationDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StorelocationDetailLogic) StorelocationDetail(in *inventory.StorelocationDetailRequest) (*inventory.StorelocationDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.StorelocationDetailResponse{}, nil
}
