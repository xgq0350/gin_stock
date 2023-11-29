package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProchecklistCheckSingleExLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProchecklistCheckSingleExLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProchecklistCheckSingleExLogic {
	return &ProchecklistCheckSingleExLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProchecklistCheckSingleExLogic) ProchecklistCheckSingleEx(in *inventory.ProchecklistCheckSingleExRequest) (*inventory.ProchecklistCheckSingleExResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.ProchecklistCheckSingleExResponse{}, nil
}
