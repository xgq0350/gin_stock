package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type StorelocationListExLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStorelocationListExLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StorelocationListExLogic {
	return &StorelocationListExLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StorelocationListExLogic) StorelocationListEx(in *inventory.StorelocationListRequest) (*inventory.StorelocationListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.StorelocationListResponse{}, nil
}
