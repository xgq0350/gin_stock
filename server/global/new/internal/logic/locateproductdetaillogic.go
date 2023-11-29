package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type LocateproductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLocateproductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LocateproductDetailLogic {
	return &LocateproductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LocateproductDetailLogic) LocateproductDetail(in *inventory.NormalDetailRequest) (*inventory.LocateproductDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.LocateproductDetailResponse{}, nil
}
