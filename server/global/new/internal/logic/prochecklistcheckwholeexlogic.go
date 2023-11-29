package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProchecklistCheckWholeExLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProchecklistCheckWholeExLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProchecklistCheckWholeExLogic {
	return &ProchecklistCheckWholeExLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProchecklistCheckWholeExLogic) ProchecklistCheckWholeEx(in *inventory.ProchecklistCheckSingleExRequest) (*inventory.ProchecklistCheckSingleExResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.ProchecklistCheckSingleExResponse{}, nil
}
