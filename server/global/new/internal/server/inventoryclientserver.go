// Code generated by goctl. DO NOT EDIT.
// Source: inventory.proto

package server

import (
	"context"

	"mall/service/inventory/rpc/new/internal/logic"
	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"
)

type InventoryClientServer struct {
	svcCtx *svc.ServiceContext
	inventory.UnimplementedInventoryClientServer
}

func NewInventoryClientServer(svcCtx *svc.ServiceContext) *InventoryClientServer {
	return &InventoryClientServer{
		svcCtx: svcCtx,
	}
}

func (s *InventoryClientServer) InventoryDelete(ctx context.Context, in *inventory.NormalDeleteRequest) (*inventory.NormalResponse, error) {
	l := logic.NewInventoryDeleteLogic(ctx, s.svcCtx)
	return l.InventoryDelete(in)
}

func (s *InventoryClientServer) InventoryCreate(ctx context.Context, in *inventory.InventoryCreateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewInventoryCreateLogic(ctx, s.svcCtx)
	return l.InventoryCreate(in)
}

func (s *InventoryClientServer) InventoryUpdate(ctx context.Context, in *inventory.InventoryUpdateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewInventoryUpdateLogic(ctx, s.svcCtx)
	return l.InventoryUpdate(in)
}

func (s *InventoryClientServer) InventoryDetail(ctx context.Context, in *inventory.NormalDetailRequest) (*inventory.InventoryDetailResponse, error) {
	l := logic.NewInventoryDetailLogic(ctx, s.svcCtx)
	return l.InventoryDetail(in)
}

func (s *InventoryClientServer) InventoryOne(ctx context.Context, in *inventory.NormalDetailRequest) (*inventory.InventoryDetailResponse, error) {
	l := logic.NewInventoryOneLogic(ctx, s.svcCtx)
	return l.InventoryOne(in)
}

func (s *InventoryClientServer) InventoryList(ctx context.Context, in *inventory.NormalListRequest) (*inventory.InventoryListResponse, error) {
	l := logic.NewInventoryListLogic(ctx, s.svcCtx)
	return l.InventoryList(in)
}

func (s *InventoryClientServer) ProchecklistDelete(ctx context.Context, in *inventory.NormalDeleteRequest) (*inventory.NormalResponse, error) {
	l := logic.NewProchecklistDeleteLogic(ctx, s.svcCtx)
	return l.ProchecklistDelete(in)
}

func (s *InventoryClientServer) ProchecklistCheckSingle(ctx context.Context, in *inventory.ProchecklistCheckSingleRequest) (*inventory.ProchecklistCheckSingleResponse, error) {
	l := logic.NewProchecklistCheckSingleLogic(ctx, s.svcCtx)
	return l.ProchecklistCheckSingle(in)
}

func (s *InventoryClientServer) ProchecklistCheckWhole(ctx context.Context, in *inventory.ProchecklistCheckSingleRequest) (*inventory.ProchecklistCheckSingleResponse, error) {
	l := logic.NewProchecklistCheckWholeLogic(ctx, s.svcCtx)
	return l.ProchecklistCheckWhole(in)
}

func (s *InventoryClientServer) ProchecklistCheckSingleEx(ctx context.Context, in *inventory.ProchecklistCheckSingleExRequest) (*inventory.ProchecklistCheckSingleExResponse, error) {
	l := logic.NewProchecklistCheckSingleExLogic(ctx, s.svcCtx)
	return l.ProchecklistCheckSingleEx(in)
}

func (s *InventoryClientServer) ProchecklistCheckWholeEx(ctx context.Context, in *inventory.ProchecklistCheckSingleExRequest) (*inventory.ProchecklistCheckSingleExResponse, error) {
	l := logic.NewProchecklistCheckWholeExLogic(ctx, s.svcCtx)
	return l.ProchecklistCheckWholeEx(in)
}

func (s *InventoryClientServer) ProchecklistCreate(ctx context.Context, in *inventory.ProchecklistCreateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewProchecklistCreateLogic(ctx, s.svcCtx)
	return l.ProchecklistCreate(in)
}

func (s *InventoryClientServer) ProchecklistUpdate(ctx context.Context, in *inventory.ProchecklistUpdateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewProchecklistUpdateLogic(ctx, s.svcCtx)
	return l.ProchecklistUpdate(in)
}

func (s *InventoryClientServer) ProchecklistDetail(ctx context.Context, in *inventory.NormalDetailRequest) (*inventory.ProchecklistDetailResponse, error) {
	l := logic.NewProchecklistDetailLogic(ctx, s.svcCtx)
	return l.ProchecklistDetail(in)
}

func (s *InventoryClientServer) ProchecklistList(ctx context.Context, in *inventory.NormalListRequest) (*inventory.ProchecklistListResponse, error) {
	l := logic.NewProchecklistListLogic(ctx, s.svcCtx)
	return l.ProchecklistList(in)
}

func (s *InventoryClientServer) ProchecklistListEx(ctx context.Context, in *inventory.NormalListRequest) (*inventory.ProchecklistListResponse, error) {
	l := logic.NewProchecklistListExLogic(ctx, s.svcCtx)
	return l.ProchecklistListEx(in)
}

func (s *InventoryClientServer) OperproductDelete(ctx context.Context, in *inventory.NormalDeleteRequest) (*inventory.NormalResponse, error) {
	l := logic.NewOperproductDeleteLogic(ctx, s.svcCtx)
	return l.OperproductDelete(in)
}

func (s *InventoryClientServer) OperproductCreate(ctx context.Context, in *inventory.OperproductCreateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewOperproductCreateLogic(ctx, s.svcCtx)
	return l.OperproductCreate(in)
}

func (s *InventoryClientServer) OperproductUpdate(ctx context.Context, in *inventory.OperproductUpdateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewOperproductUpdateLogic(ctx, s.svcCtx)
	return l.OperproductUpdate(in)
}

func (s *InventoryClientServer) OperproductDetail(ctx context.Context, in *inventory.NormalDetailRequest) (*inventory.OperproductDetailResponse, error) {
	l := logic.NewOperproductDetailLogic(ctx, s.svcCtx)
	return l.OperproductDetail(in)
}

func (s *InventoryClientServer) OperproductList(ctx context.Context, in *inventory.NormalListRequest) (*inventory.OperproductListResponse, error) {
	l := logic.NewOperproductListLogic(ctx, s.svcCtx)
	return l.OperproductList(in)
}

func (s *InventoryClientServer) LocateproductDelete(ctx context.Context, in *inventory.NormalDeleteRequest) (*inventory.NormalResponse, error) {
	l := logic.NewLocateproductDeleteLogic(ctx, s.svcCtx)
	return l.LocateproductDelete(in)
}

func (s *InventoryClientServer) LocateproductCreate(ctx context.Context, in *inventory.LocateproductCreateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewLocateproductCreateLogic(ctx, s.svcCtx)
	return l.LocateproductCreate(in)
}

func (s *InventoryClientServer) LocateproductUpdate(ctx context.Context, in *inventory.LocateproductUpdateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewLocateproductUpdateLogic(ctx, s.svcCtx)
	return l.LocateproductUpdate(in)
}

func (s *InventoryClientServer) LocateproductDetail(ctx context.Context, in *inventory.NormalDetailRequest) (*inventory.LocateproductDetailResponse, error) {
	l := logic.NewLocateproductDetailLogic(ctx, s.svcCtx)
	return l.LocateproductDetail(in)
}

func (s *InventoryClientServer) LocateproductList(ctx context.Context, in *inventory.NormalListRequest) (*inventory.LocateproductListResponse, error) {
	l := logic.NewLocateproductListLogic(ctx, s.svcCtx)
	return l.LocateproductList(in)
}

func (s *InventoryClientServer) OrderDelete(ctx context.Context, in *inventory.NormalDeleteRequest) (*inventory.NormalResponse, error) {
	l := logic.NewOrderDeleteLogic(ctx, s.svcCtx)
	return l.OrderDelete(in)
}

func (s *InventoryClientServer) OrderCreate(ctx context.Context, in *inventory.OrderCreateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewOrderCreateLogic(ctx, s.svcCtx)
	return l.OrderCreate(in)
}

func (s *InventoryClientServer) OrderUpdate(ctx context.Context, in *inventory.OrderUpdateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewOrderUpdateLogic(ctx, s.svcCtx)
	return l.OrderUpdate(in)
}

func (s *InventoryClientServer) OrderDetail(ctx context.Context, in *inventory.NormalDetailRequest) (*inventory.OrderDetailResponse, error) {
	l := logic.NewOrderDetailLogic(ctx, s.svcCtx)
	return l.OrderDetail(in)
}

func (s *InventoryClientServer) OrderGetLastDetail(ctx context.Context, in *inventory.NormalDetailRequest) (*inventory.OrderDetailResponse, error) {
	l := logic.NewOrderGetLastDetailLogic(ctx, s.svcCtx)
	return l.OrderGetLastDetail(in)
}

func (s *InventoryClientServer) OrderList(ctx context.Context, in *inventory.NormalStatusListRequest) (*inventory.OrderListResponse, error) {
	l := logic.NewOrderListLogic(ctx, s.svcCtx)
	return l.OrderList(in)
}

func (s *InventoryClientServer) OrderClone(ctx context.Context, in *inventory.NormalIDListRequest) (*inventory.NormalResponse, error) {
	l := logic.NewOrderCloneLogic(ctx, s.svcCtx)
	return l.OrderClone(in)
}

func (s *InventoryClientServer) OrderaddressDelete(ctx context.Context, in *inventory.NormalDeleteRequest) (*inventory.NormalResponse, error) {
	l := logic.NewOrderaddressDeleteLogic(ctx, s.svcCtx)
	return l.OrderaddressDelete(in)
}

func (s *InventoryClientServer) OrderaddressCreate(ctx context.Context, in *inventory.OrderaddressCreateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewOrderaddressCreateLogic(ctx, s.svcCtx)
	return l.OrderaddressCreate(in)
}

func (s *InventoryClientServer) OrderaddressUpdate(ctx context.Context, in *inventory.OrderaddressUpdateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewOrderaddressUpdateLogic(ctx, s.svcCtx)
	return l.OrderaddressUpdate(in)
}

func (s *InventoryClientServer) OrderaddressDetail(ctx context.Context, in *inventory.NormalDetailRequest) (*inventory.OrderaddressDetailResponse, error) {
	l := logic.NewOrderaddressDetailLogic(ctx, s.svcCtx)
	return l.OrderaddressDetail(in)
}

func (s *InventoryClientServer) OrderaddressList(ctx context.Context, in *inventory.NormalListRequest) (*inventory.OrderaddressListResponse, error) {
	l := logic.NewOrderaddressListLogic(ctx, s.svcCtx)
	return l.OrderaddressList(in)
}

func (s *InventoryClientServer) OrderproductDelete(ctx context.Context, in *inventory.NormalDeleteRequest) (*inventory.NormalResponse, error) {
	l := logic.NewOrderproductDeleteLogic(ctx, s.svcCtx)
	return l.OrderproductDelete(in)
}

func (s *InventoryClientServer) OrderproductCreate(ctx context.Context, in *inventory.OrderproductCreateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewOrderproductCreateLogic(ctx, s.svcCtx)
	return l.OrderproductCreate(in)
}

func (s *InventoryClientServer) OrderproductUpdate(ctx context.Context, in *inventory.OrderproductUpdateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewOrderproductUpdateLogic(ctx, s.svcCtx)
	return l.OrderproductUpdate(in)
}

func (s *InventoryClientServer) OrderproductDetail(ctx context.Context, in *inventory.NormalDetailRequest) (*inventory.OrderproductDetailResponse, error) {
	l := logic.NewOrderproductDetailLogic(ctx, s.svcCtx)
	return l.OrderproductDetail(in)
}

func (s *InventoryClientServer) OrderproductList(ctx context.Context, in *inventory.NormalStatusListRequest) (*inventory.OrderproductListResponse, error) {
	l := logic.NewOrderproductListLogic(ctx, s.svcCtx)
	return l.OrderproductList(in)
}

func (s *InventoryClientServer) WmsOrderproductList(ctx context.Context, in *inventory.NormalIDListRequest) (*inventory.OrderproductListResponse, error) {
	l := logic.NewWmsOrderproductListLogic(ctx, s.svcCtx)
	return l.WmsOrderproductList(in)
}

func (s *InventoryClientServer) WmsOrderUserproductList(ctx context.Context, in *inventory.NormalIDListRequest) (*inventory.OrderproductListResponse, error) {
	l := logic.NewWmsOrderUserproductListLogic(ctx, s.svcCtx)
	return l.WmsOrderUserproductList(in)
}

func (s *InventoryClientServer) WmsPurchaseproductList(ctx context.Context, in *inventory.NormalIDListRequest) (*inventory.ProchecklistListResponse, error) {
	l := logic.NewWmsPurchaseproductListLogic(ctx, s.svcCtx)
	return l.WmsPurchaseproductList(in)
}

func (s *InventoryClientServer) WmsPurchaseproductListEx(ctx context.Context, in *inventory.NormalIDListRequest) (*inventory.ProchecklistListResponse, error) {
	l := logic.NewWmsPurchaseproductListExLogic(ctx, s.svcCtx)
	return l.WmsPurchaseproductListEx(in)
}

func (s *InventoryClientServer) WmsPurchaseUserproductList(ctx context.Context, in *inventory.NormalIDListRequest) (*inventory.ProchecklistListResponse, error) {
	l := logic.NewWmsPurchaseUserproductListLogic(ctx, s.svcCtx)
	return l.WmsPurchaseUserproductList(in)
}

func (s *InventoryClientServer) WmsOrderUserproductListPdf(ctx context.Context, in *inventory.NormalIDListRequest) (*inventory.NormalResponsePdf, error) {
	l := logic.NewWmsOrderUserproductListPdfLogic(ctx, s.svcCtx)
	return l.WmsOrderUserproductListPdf(in)
}

func (s *InventoryClientServer) WmsPurchaseListPdf(ctx context.Context, in *inventory.NormalIDListRequest) (*inventory.NormalResponsePdf, error) {
	l := logic.NewWmsPurchaseListPdfLogic(ctx, s.svcCtx)
	return l.WmsPurchaseListPdf(in)
}

func (s *InventoryClientServer) WmsProcheckPurchaseProductListEx(ctx context.Context, in *inventory.NormalIDListRequest) (*inventory.ProchecklistListResponse, error) {
	l := logic.NewWmsProcheckPurchaseProductListExLogic(ctx, s.svcCtx)
	return l.WmsProcheckPurchaseProductListEx(in)
}

func (s *InventoryClientServer) LogisticsDelete(ctx context.Context, in *inventory.NormalDeleteRequest) (*inventory.NormalResponse, error) {
	l := logic.NewLogisticsDeleteLogic(ctx, s.svcCtx)
	return l.LogisticsDelete(in)
}

func (s *InventoryClientServer) LogisticsCreate(ctx context.Context, in *inventory.LogisticsCreateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewLogisticsCreateLogic(ctx, s.svcCtx)
	return l.LogisticsCreate(in)
}

func (s *InventoryClientServer) LogisticsUpdate(ctx context.Context, in *inventory.LogisticsUpdateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewLogisticsUpdateLogic(ctx, s.svcCtx)
	return l.LogisticsUpdate(in)
}

func (s *InventoryClientServer) LogisticsDetail(ctx context.Context, in *inventory.NormalDetailRequest) (*inventory.LogisticsDetailResponse, error) {
	l := logic.NewLogisticsDetailLogic(ctx, s.svcCtx)
	return l.LogisticsDetail(in)
}

func (s *InventoryClientServer) LogisticsList(ctx context.Context, in *inventory.NormalListRequest) (*inventory.LogisticsListResponse, error) {
	l := logic.NewLogisticsListLogic(ctx, s.svcCtx)
	return l.LogisticsList(in)
}

func (s *InventoryClientServer) WorkerorderDelete(ctx context.Context, in *inventory.NormalDeleteRequest) (*inventory.NormalResponse, error) {
	l := logic.NewWorkerorderDeleteLogic(ctx, s.svcCtx)
	return l.WorkerorderDelete(in)
}

func (s *InventoryClientServer) WorkerorderCreate(ctx context.Context, in *inventory.WorkerorderCreateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewWorkerorderCreateLogic(ctx, s.svcCtx)
	return l.WorkerorderCreate(in)
}

func (s *InventoryClientServer) WorkerorderUpdate(ctx context.Context, in *inventory.WorkerorderUpdateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewWorkerorderUpdateLogic(ctx, s.svcCtx)
	return l.WorkerorderUpdate(in)
}

func (s *InventoryClientServer) WorkerorderDetail(ctx context.Context, in *inventory.NormalDetailRequest) (*inventory.WorkerorderDetailResponse, error) {
	l := logic.NewWorkerorderDetailLogic(ctx, s.svcCtx)
	return l.WorkerorderDetail(in)
}

func (s *InventoryClientServer) WorkerorderList(ctx context.Context, in *inventory.NormalStatusListRequest) (*inventory.WorkerorderListResponse, error) {
	l := logic.NewWorkerorderListLogic(ctx, s.svcCtx)
	return l.WorkerorderList(in)
}

func (s *InventoryClientServer) WorkerorderOrderlist(ctx context.Context, in *inventory.NormalStatusListRequest) (*inventory.WorkerorderListResponse, error) {
	l := logic.NewWorkerorderOrderlistLogic(ctx, s.svcCtx)
	return l.WorkerorderOrderlist(in)
}

func (s *InventoryClientServer) WaveCreate(ctx context.Context, in *inventory.WaveCreateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewWaveCreateLogic(ctx, s.svcCtx)
	return l.WaveCreate(in)
}

func (s *InventoryClientServer) WaveOperator(ctx context.Context, in *inventory.NormalStatusListRequest) (*inventory.NormalResponse, error) {
	l := logic.NewWaveOperatorLogic(ctx, s.svcCtx)
	return l.WaveOperator(in)
}

func (s *InventoryClientServer) WaveProductUpdateStatus(ctx context.Context, in *inventory.NormalStatusListRequest) (*inventory.NormalResponse, error) {
	l := logic.NewWaveProductUpdateStatusLogic(ctx, s.svcCtx)
	return l.WaveProductUpdateStatus(in)
}

func (s *InventoryClientServer) Wavelist(ctx context.Context, in *inventory.NormalListRequest) (*inventory.WaveListResponse, error) {
	l := logic.NewWavelistLogic(ctx, s.svcCtx)
	return l.Wavelist(in)
}

func (s *InventoryClientServer) WaveProductlist(ctx context.Context, in *inventory.NormalIDListRequest) (*inventory.WaveProductListResponse, error) {
	l := logic.NewWaveProductlistLogic(ctx, s.svcCtx)
	return l.WaveProductlist(in)
}

func (s *InventoryClientServer) WaveOrderlist(ctx context.Context, in *inventory.NormalIDListRequest) (*inventory.WorkerorderListResponse, error) {
	l := logic.NewWaveOrderlistLogic(ctx, s.svcCtx)
	return l.WaveOrderlist(in)
}

func (s *InventoryClientServer) WaveOrderProductlist(ctx context.Context, in *inventory.NormalIDListRequest) (*inventory.OrderproductListResponse, error) {
	l := logic.NewWaveOrderProductlistLogic(ctx, s.svcCtx)
	return l.WaveOrderProductlist(in)
}

func (s *InventoryClientServer) WaveConfirmOutProduct(ctx context.Context, in *inventory.ConfirmOutProductRequest) (*inventory.NormalResponse, error) {
	l := logic.NewWaveConfirmOutProductLogic(ctx, s.svcCtx)
	return l.WaveConfirmOutProduct(in)
}

func (s *InventoryClientServer) WaveConfirmOutProductEx(ctx context.Context, in *inventory.ConfirmOutProductRequest) (*inventory.NormalResponse, error) {
	l := logic.NewWaveConfirmOutProductExLogic(ctx, s.svcCtx)
	return l.WaveConfirmOutProductEx(in)
}

func (s *InventoryClientServer) WaveOrderproductListPdf(ctx context.Context, in *inventory.NormalIDListRequest) (*inventory.NormalResponsePdf, error) {
	l := logic.NewWaveOrderproductListPdfLogic(ctx, s.svcCtx)
	return l.WaveOrderproductListPdf(in)
}

func (s *InventoryClientServer) Inoutloglist(ctx context.Context, in *inventory.NormalListRequest) (*inventory.InoutlogListResponse, error) {
	l := logic.NewInoutloglistLogic(ctx, s.svcCtx)
	return l.Inoutloglist(in)
}

func (s *InventoryClientServer) Locationfeelist(ctx context.Context, in *inventory.NormalListRequest) (*inventory.LocationfeeListResponse, error) {
	l := logic.NewLocationfeelistLogic(ctx, s.svcCtx)
	return l.Locationfeelist(in)
}

func (s *InventoryClientServer) LocationfeeDayQeryList(ctx context.Context, in *inventory.LocationfeeDayqueryRequest) (*inventory.LocationfeeDayqueryResponse, error) {
	l := logic.NewLocationfeeDayQeryListLogic(ctx, s.svcCtx)
	return l.LocationfeeDayQeryList(in)
}

func (s *InventoryClientServer) ConfirmRecvProduct(ctx context.Context, in *inventory.ConfirmRecvProductRequest) (*inventory.NormalResponse, error) {
	l := logic.NewConfirmRecvProductLogic(ctx, s.svcCtx)
	return l.ConfirmRecvProduct(in)
}

func (s *InventoryClientServer) ConfirmOutProduct(ctx context.Context, in *inventory.ConfirmOutProductRequest) (*inventory.NormalResponse, error) {
	l := logic.NewConfirmOutProductLogic(ctx, s.svcCtx)
	return l.ConfirmOutProduct(in)
}

func (s *InventoryClientServer) CustwarehouseDelete(ctx context.Context, in *inventory.NormalDeleteRequest) (*inventory.NormalResponse, error) {
	l := logic.NewCustwarehouseDeleteLogic(ctx, s.svcCtx)
	return l.CustwarehouseDelete(in)
}

func (s *InventoryClientServer) CustwarehouseCreate(ctx context.Context, in *inventory.CustwarehouseCreateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewCustwarehouseCreateLogic(ctx, s.svcCtx)
	return l.CustwarehouseCreate(in)
}

func (s *InventoryClientServer) CustwarehouseUpdate(ctx context.Context, in *inventory.CustwarehouseUpdateRequest) (*inventory.NormalResponse, error) {
	l := logic.NewCustwarehouseUpdateLogic(ctx, s.svcCtx)
	return l.CustwarehouseUpdate(in)
}

func (s *InventoryClientServer) CustwarehouseDetail(ctx context.Context, in *inventory.NormalDetailRequest) (*inventory.CustwarehouseDetailResponse, error) {
	l := logic.NewCustwarehouseDetailLogic(ctx, s.svcCtx)
	return l.CustwarehouseDetail(in)
}

func (s *InventoryClientServer) CustwarehouseList(ctx context.Context, in *inventory.NormalListRequest) (*inventory.CustwarehouseListResponse, error) {
	l := logic.NewCustwarehouseListLogic(ctx, s.svcCtx)
	return l.CustwarehouseList(in)
}

func (s *InventoryClientServer) CustwarehouseListEx(ctx context.Context, in *inventory.NormalListRequest) (*inventory.CustwarehouseListResponse, error) {
	l := logic.NewCustwarehouseListExLogic(ctx, s.svcCtx)
	return l.CustwarehouseListEx(in)
}

func (s *InventoryClientServer) AddressList(ctx context.Context, in *inventory.AddressListRequest) (*inventory.AddressListResponse, error) {
	l := logic.NewAddressListLogic(ctx, s.svcCtx)
	return l.AddressList(in)
}

func (s *InventoryClientServer) AddressDetail(ctx context.Context, in *inventory.AddressDetailRequest) (*inventory.AddressDetailResponse, error) {
	l := logic.NewAddressDetailLogic(ctx, s.svcCtx)
	return l.AddressDetail(in)
}

func (s *InventoryClientServer) AddressCreate(ctx context.Context, in *inventory.AddressCreateRequest) (*inventory.AddressCreateResponse, error) {
	l := logic.NewAddressCreateLogic(ctx, s.svcCtx)
	return l.AddressCreate(in)
}

func (s *InventoryClientServer) AddressUpdate(ctx context.Context, in *inventory.AddressUpdateRequest) (*inventory.AddressUpdateResponse, error) {
	l := logic.NewAddressUpdateLogic(ctx, s.svcCtx)
	return l.AddressUpdate(in)
}

func (s *InventoryClientServer) AddressDelete(ctx context.Context, in *inventory.AddressRemoveRequest) (*inventory.AddressRemoveResponse, error) {
	l := logic.NewAddressDeleteLogic(ctx, s.svcCtx)
	return l.AddressDelete(in)
}

func (s *InventoryClientServer) ProductList(ctx context.Context, in *inventory.ProductListRequest) (*inventory.ProductListResponse, error) {
	l := logic.NewProductListLogic(ctx, s.svcCtx)
	return l.ProductList(in)
}

func (s *InventoryClientServer) ProductListAdmin(ctx context.Context, in *inventory.ProductListRequest) (*inventory.ProductListResponse, error) {
	l := logic.NewProductListAdminLogic(ctx, s.svcCtx)
	return l.ProductListAdmin(in)
}

func (s *InventoryClientServer) ProductDetail(ctx context.Context, in *inventory.ProductDetailRequest) (*inventory.ProductDetailResponse, error) {
	l := logic.NewProductDetailLogic(ctx, s.svcCtx)
	return l.ProductDetail(in)
}

func (s *InventoryClientServer) ProductCreate(ctx context.Context, in *inventory.ProductCreateRequest) (*inventory.ProductCreateResponse, error) {
	l := logic.NewProductCreateLogic(ctx, s.svcCtx)
	return l.ProductCreate(in)
}

func (s *InventoryClientServer) ProductUpdate(ctx context.Context, in *inventory.ProductUpdateRequest) (*inventory.ProductUpdateResponse, error) {
	l := logic.NewProductUpdateLogic(ctx, s.svcCtx)
	return l.ProductUpdate(in)
}

func (s *InventoryClientServer) ProductDelete(ctx context.Context, in *inventory.ProductRemoveRequest) (*inventory.ProductRemoveResponse, error) {
	l := logic.NewProductDeleteLogic(ctx, s.svcCtx)
	return l.ProductDelete(in)
}

func (s *InventoryClientServer) WarehouseList(ctx context.Context, in *inventory.WarehouseListRequest) (*inventory.WarehouseListResponse, error) {
	l := logic.NewWarehouseListLogic(ctx, s.svcCtx)
	return l.WarehouseList(in)
}

func (s *InventoryClientServer) WarehouseEnable(ctx context.Context, in *inventory.NormalEnabledRequest) (*inventory.NormalResponse, error) {
	l := logic.NewWarehouseEnableLogic(ctx, s.svcCtx)
	return l.WarehouseEnable(in)
}

func (s *InventoryClientServer) WarehouseDetail(ctx context.Context, in *inventory.WarehouseDetailRequest) (*inventory.WarehouseDetailResponse, error) {
	l := logic.NewWarehouseDetailLogic(ctx, s.svcCtx)
	return l.WarehouseDetail(in)
}

func (s *InventoryClientServer) WarehouseCreate(ctx context.Context, in *inventory.WarehouseCreateRequest) (*inventory.WarehouseCreateResponse, error) {
	l := logic.NewWarehouseCreateLogic(ctx, s.svcCtx)
	return l.WarehouseCreate(in)
}

func (s *InventoryClientServer) WarehouseUpdate(ctx context.Context, in *inventory.WarehouseUpdateRequest) (*inventory.WarehouseUpdateResponse, error) {
	l := logic.NewWarehouseUpdateLogic(ctx, s.svcCtx)
	return l.WarehouseUpdate(in)
}

func (s *InventoryClientServer) WarehouseDelete(ctx context.Context, in *inventory.WarehouseRemoveRequest) (*inventory.WarehouseRemoveResponse, error) {
	l := logic.NewWarehouseDeleteLogic(ctx, s.svcCtx)
	return l.WarehouseDelete(in)
}

func (s *InventoryClientServer) StorelocationList(ctx context.Context, in *inventory.StorelocationListRequest) (*inventory.StorelocationListResponse, error) {
	l := logic.NewStorelocationListLogic(ctx, s.svcCtx)
	return l.StorelocationList(in)
}

func (s *InventoryClientServer) StorelocationListEx(ctx context.Context, in *inventory.StorelocationListRequest) (*inventory.StorelocationListResponse, error) {
	l := logic.NewStorelocationListExLogic(ctx, s.svcCtx)
	return l.StorelocationListEx(in)
}

func (s *InventoryClientServer) StorelocationDetail(ctx context.Context, in *inventory.StorelocationDetailRequest) (*inventory.StorelocationDetailResponse, error) {
	l := logic.NewStorelocationDetailLogic(ctx, s.svcCtx)
	return l.StorelocationDetail(in)
}

func (s *InventoryClientServer) StorelocationCreate(ctx context.Context, in *inventory.StorelocationCreateRequest) (*inventory.StorelocationCreateResponse, error) {
	l := logic.NewStorelocationCreateLogic(ctx, s.svcCtx)
	return l.StorelocationCreate(in)
}

func (s *InventoryClientServer) StorelocationUpdate(ctx context.Context, in *inventory.StorelocationUpdateRequest) (*inventory.StorelocationUpdateResponse, error) {
	l := logic.NewStorelocationUpdateLogic(ctx, s.svcCtx)
	return l.StorelocationUpdate(in)
}

func (s *InventoryClientServer) StorelocationDelete(ctx context.Context, in *inventory.StorelocationRemoveRequest) (*inventory.StorelocationRemoveResponse, error) {
	l := logic.NewStorelocationDeleteLogic(ctx, s.svcCtx)
	return l.StorelocationDelete(in)
}

func (s *InventoryClientServer) StorelocationMulCreate(ctx context.Context, in *inventory.StorelocationMulCreateRequest) (*inventory.StorelocationMulCreateResponse, error) {
	l := logic.NewStorelocationMulCreateLogic(ctx, s.svcCtx)
	return l.StorelocationMulCreate(in)
}

func (s *InventoryClientServer) StorelocationLocListPdf(ctx context.Context, in *inventory.NormalIDListRequest) (*inventory.NormalResponsePdf, error) {
	l := logic.NewStorelocationLocListPdfLogic(ctx, s.svcCtx)
	return l.StorelocationLocListPdf(in)
}

func (s *InventoryClientServer) StorelocationMulCreateEx(ctx context.Context, in *inventory.StorelocationMulCreateRequest) (*inventory.StorelocationMulCreateResponse, error) {
	l := logic.NewStorelocationMulCreateExLogic(ctx, s.svcCtx)
	return l.StorelocationMulCreateEx(in)
}

func (s *InventoryClientServer) StorelocationLocListPdfEx(ctx context.Context, in *inventory.NormalIDListRequest) (*inventory.NormalResponsePdf, error) {
	l := logic.NewStorelocationLocListPdfExLogic(ctx, s.svcCtx)
	return l.StorelocationLocListPdfEx(in)
}

func (s *InventoryClientServer) StorelocationListPdf(ctx context.Context, in *inventory.StorelocationListRequest) (*inventory.NormalResponsePdf, error) {
	l := logic.NewStorelocationListPdfLogic(ctx, s.svcCtx)
	return l.StorelocationListPdf(in)
}

func (s *InventoryClientServer) WmsProductList(ctx context.Context, in *inventory.WmsProductListRequest) (*inventory.WmsProductListResponse, error) {
	l := logic.NewWmsProductListLogic(ctx, s.svcCtx)
	return l.WmsProductList(in)
}

func (s *InventoryClientServer) WmsProductDetail(ctx context.Context, in *inventory.WmsProductDetailRequest) (*inventory.WmsProductDetailResponse, error) {
	l := logic.NewWmsProductDetailLogic(ctx, s.svcCtx)
	return l.WmsProductDetail(in)
}

func (s *InventoryClientServer) WmsProductCreate(ctx context.Context, in *inventory.WmsProductCreateRequest) (*inventory.WmsProductCreateResponse, error) {
	l := logic.NewWmsProductCreateLogic(ctx, s.svcCtx)
	return l.WmsProductCreate(in)
}

func (s *InventoryClientServer) WmsProductUpdate(ctx context.Context, in *inventory.WmsProductUpdateRequest) (*inventory.WmsProductUpdateResponse, error) {
	l := logic.NewWmsProductUpdateLogic(ctx, s.svcCtx)
	return l.WmsProductUpdate(in)
}

func (s *InventoryClientServer) WmsProductDelete(ctx context.Context, in *inventory.WmsProductRemoveRequest) (*inventory.WmsProductRemoveResponse, error) {
	l := logic.NewWmsProductDeleteLogic(ctx, s.svcCtx)
	return l.WmsProductDelete(in)
}

func (s *InventoryClientServer) RepairList(ctx context.Context, in *inventory.RepairListRequest) (*inventory.RepairListResponse, error) {
	l := logic.NewRepairListLogic(ctx, s.svcCtx)
	return l.RepairList(in)
}

func (s *InventoryClientServer) RepairListEx(ctx context.Context, in *inventory.RepairListRequest) (*inventory.RepairListResponse, error) {
	l := logic.NewRepairListExLogic(ctx, s.svcCtx)
	return l.RepairListEx(in)
}

func (s *InventoryClientServer) RepairDetail(ctx context.Context, in *inventory.RepairDetailRequest) (*inventory.RepairDetailResponse, error) {
	l := logic.NewRepairDetailLogic(ctx, s.svcCtx)
	return l.RepairDetail(in)
}

func (s *InventoryClientServer) RepairCreate(ctx context.Context, in *inventory.RepairCreateRequest) (*inventory.RepairCreateResponse, error) {
	l := logic.NewRepairCreateLogic(ctx, s.svcCtx)
	return l.RepairCreate(in)
}

func (s *InventoryClientServer) RepairUpdate(ctx context.Context, in *inventory.RepairUpdateRequest) (*inventory.RepairUpdateResponse, error) {
	l := logic.NewRepairUpdateLogic(ctx, s.svcCtx)
	return l.RepairUpdate(in)
}

func (s *InventoryClientServer) RepairDelete(ctx context.Context, in *inventory.RepairRemoveRequest) (*inventory.RepairRemoveResponse, error) {
	l := logic.NewRepairDeleteLogic(ctx, s.svcCtx)
	return l.RepairDelete(in)
}

func (s *InventoryClientServer) CustProductStoreDetail(ctx context.Context, in *inventory.NormalDetailRequest) (*inventory.ProductlocationDetailResponse, error) {
	l := logic.NewCustProductStoreDetailLogic(ctx, s.svcCtx)
	return l.CustProductStoreDetail(in)
}

func (s *InventoryClientServer) CustProductStoreList(ctx context.Context, in *inventory.NormalListRequest) (*inventory.ProductlocationListResponse, error) {
	l := logic.NewCustProductStoreListLogic(ctx, s.svcCtx)
	return l.CustProductStoreList(in)
}

func (s *InventoryClientServer) CustProductStoreListEx(ctx context.Context, in *inventory.NormalListRequest) (*inventory.ProductlocationListResponse, error) {
	l := logic.NewCustProductStoreListExLogic(ctx, s.svcCtx)
	return l.CustProductStoreListEx(in)
}

func (s *InventoryClientServer) CustProductLocationDetail(ctx context.Context, in *inventory.NormalDetailRequest) (*inventory.ProductlocationDetailResponse, error) {
	l := logic.NewCustProductLocationDetailLogic(ctx, s.svcCtx)
	return l.CustProductLocationDetail(in)
}

func (s *InventoryClientServer) CustProductLocationList(ctx context.Context, in *inventory.NormalListRequest) (*inventory.ProductlocationListResponse, error) {
	l := logic.NewCustProductLocationListLogic(ctx, s.svcCtx)
	return l.CustProductLocationList(in)
}

func (s *InventoryClientServer) CustProductLocationListEx(ctx context.Context, in *inventory.NormalListRequest) (*inventory.ProductlocationListResponse, error) {
	l := logic.NewCustProductLocationListExLogic(ctx, s.svcCtx)
	return l.CustProductLocationListEx(in)
}

func (s *InventoryClientServer) PalletPrintList(ctx context.Context, in *inventory.NormalListRequest) (*inventory.PalletPrintListResponse, error) {
	l := logic.NewPalletPrintListLogic(ctx, s.svcCtx)
	return l.PalletPrintList(in)
}

func (s *InventoryClientServer) PalletPrintDetail(ctx context.Context, in *inventory.PalletPrintRequest) (*inventory.PalletPrintListResponse, error) {
	l := logic.NewPalletPrintDetailLogic(ctx, s.svcCtx)
	return l.PalletPrintDetail(in)
}

func (s *InventoryClientServer) PalletPrint(ctx context.Context, in *inventory.PalletPrintRequest) (*inventory.PalletPrintResponse, error) {
	l := logic.NewPalletPrintLogic(ctx, s.svcCtx)
	return l.PalletPrint(in)
}

func (s *InventoryClientServer) PalletPrintUpdate(ctx context.Context, in *inventory.PalletPrintRequest) (*inventory.PalletPrintListResponse, error) {
	l := logic.NewPalletPrintUpdateLogic(ctx, s.svcCtx)
	return l.PalletPrintUpdate(in)
}

func (s *InventoryClientServer) ProchecklistUpdatePallet(ctx context.Context, in *inventory.PalletPrintRequest) (*inventory.ProchecklistUpdatePalletResponse, error) {
	l := logic.NewProchecklistUpdatePalletLogic(ctx, s.svcCtx)
	return l.ProchecklistUpdatePallet(in)
}
