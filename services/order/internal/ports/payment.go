package ports

import "github.com/omareloui/gRPCMicroservicesInGo/services/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
