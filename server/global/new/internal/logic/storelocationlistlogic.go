package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type StorelocationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStorelocationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StorelocationListLogic {
	return &StorelocationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StorelocationListLogic) StorelocationList(in *inventory.StorelocationListRequest) (*inventory.StorelocationListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.StorelocationListResponse{}, nil
}
