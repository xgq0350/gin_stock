package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type StorelocationMulCreateExLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStorelocationMulCreateExLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StorelocationMulCreateExLogic {
	return &StorelocationMulCreateExLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StorelocationMulCreateExLogic) StorelocationMulCreateEx(in *inventory.StorelocationMulCreateRequest) (*inventory.StorelocationMulCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.StorelocationMulCreateResponse{}, nil
}
