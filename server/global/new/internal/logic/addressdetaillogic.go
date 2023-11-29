package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddressDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressDetailLogic {
	return &AddressDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddressDetailLogic) AddressDetail(in *inventory.AddressDetailRequest) (*inventory.AddressDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.AddressDetailResponse{}, nil
}
