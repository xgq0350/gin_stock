package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type StorelocationMulCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStorelocationMulCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StorelocationMulCreateLogic {
	return &StorelocationMulCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StorelocationMulCreateLogic) StorelocationMulCreate(in *inventory.StorelocationMulCreateRequest) (*inventory.StorelocationMulCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.StorelocationMulCreateResponse{}, nil
}
