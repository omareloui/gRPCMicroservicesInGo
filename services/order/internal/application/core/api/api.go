package api

import (
	"github.com/omareloui/gRPCMicroservicesInGo/services/order/internal/application/core/domain"
	"github.com/omareloui/gRPCMicroservicesInGo/services/order/internal/ports"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{db: db, payment: payment}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	err = a.payment.Charge(&order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (a Application) GetOrder(id uint) (domain.Order, error) {
	order, err := a.db.Get(id)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
