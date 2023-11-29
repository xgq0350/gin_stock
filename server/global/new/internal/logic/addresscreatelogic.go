package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddressCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressCreateLogic {
	return &AddressCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddressCreateLogic) AddressCreate(in *inventory.AddressCreateRequest) (*inventory.AddressCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.AddressCreateResponse{}, nil
}
