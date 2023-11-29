package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProchecklistListExLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProchecklistListExLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProchecklistListExLogic {
	return &ProchecklistListExLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProchecklistListExLogic) ProchecklistListEx(in *inventory.NormalListRequest) (*inventory.ProchecklistListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.ProchecklistListResponse{}, nil
}
