package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperproductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOperproductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperproductDetailLogic {
	return &OperproductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OperproductDetailLogic) OperproductDetail(in *inventory.NormalDetailRequest) (*inventory.OperproductDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.OperproductDetailResponse{}, nil
}
