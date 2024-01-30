package grpc

import (
	"context"

	"github.com/omareloui/gRPCMicroservicesInGo/proto/go/order"
	"github.com/omareloui/gRPCMicroservicesInGo/services/order/internal/application/core/domain"
)

func (a *Adapter) Create(ctx context.Context, request *order.CreateRequest) (*order.CreateResponse, error) {
	var orderItems []domain.OrderItem
	for _, item := range request.Items {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}
	newOrder := domain.NewOrder(uint(request.UserId), orderItems)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}
	return &order.CreateResponse{OrderId: int64(result.ID)}, nil
}

func (a *Adapter) Get(ctx context.Context, request *order.GetRequest) (*order.OrderM, error) {
	o, err := a.api.GetOrder(uint(request.OrderId))
	if err != nil {
		return nil, err
	}
	var orderItems []*order.Item
	for _, item := range o.OrderItems {
		orderItems = append(orderItems, &order.Item{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}
	return &order.OrderM{
		Id:         int64(o.ID),
		UserId:     int64(o.CustomerID),
		TotalPrice: o.TotalPrice(),
		Items:      orderItems,
	}, nil
}
