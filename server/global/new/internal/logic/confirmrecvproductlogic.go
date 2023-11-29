package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfirmRecvProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfirmRecvProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfirmRecvProductLogic {
	return &ConfirmRecvProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfirmRecvProductLogic) ConfirmRecvProduct(in *inventory.ConfirmRecvProductRequest) (*inventory.NormalResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.NormalResponse{}, nil
}
