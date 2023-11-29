package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogisticsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogisticsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogisticsListLogic {
	return &LogisticsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogisticsListLogic) LogisticsList(in *inventory.NormalListRequest) (*inventory.LogisticsListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.LogisticsListResponse{}, nil
}
