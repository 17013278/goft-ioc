package Config

import (
	"github.com/17013278/goft-ioc/examples/services"
)

type ServiceConfig struct {
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}
func (this *ServiceConfig) OrderService() *services.OrderService {
	return services.NewOrderService()
}
