package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProchecklistDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProchecklistDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProchecklistDetailLogic {
	return &ProchecklistDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProchecklistDetailLogic) ProchecklistDetail(in *inventory.NormalDetailRequest) (*inventory.ProchecklistDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.ProchecklistDetailResponse{}, nil
}
