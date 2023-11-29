package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/controller"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/stock"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	ExampleServiceGroup    example.ServiceGroup
	StockServiceGroup      stock.ServiceGroup
	StockServiceController controller.ControllerGroup
}

var ServiceGroupApp = new(ServiceGroup)
