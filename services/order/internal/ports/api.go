package ports

import "github.com/omareloui/gRPCMicroservicesInGo/services/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
	GetOrder(id uint) (domain.Order, error)
}
