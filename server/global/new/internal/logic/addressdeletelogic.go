package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressDeleteLogic {
	return &AddressDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddressDeleteLogic) AddressDelete(in *inventory.AddressRemoveRequest) (*inventory.AddressRemoveResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.AddressRemoveResponse{}, nil
}
