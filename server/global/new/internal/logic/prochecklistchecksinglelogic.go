package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProchecklistCheckSingleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProchecklistCheckSingleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProchecklistCheckSingleLogic {
	return &ProchecklistCheckSingleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProchecklistCheckSingleLogic) ProchecklistCheckSingle(in *inventory.ProchecklistCheckSingleRequest) (*inventory.ProchecklistCheckSingleResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.ProchecklistCheckSingleResponse{}, nil
}
