package db

import (
	"fmt"

	"github.com/omareloui/gRPCMicroservicesInGo/services/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Adapter struct {
	db *gorm.DB
}

type Order struct {
	gorm.Model
	CustomerID uint
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    uint32
	OrderID     uint
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection error: %v", err)
	}

	err = db.AutoMigrate(&Order{}, OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (a Adapter) Get(id uint) (domain.Order, error) {
	var orderEntity Order
	res := a.db.First(&orderEntity, id)

	var orderItems []domain.OrderItem
	for _, orderItem := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	order := domain.Order{
		ID:         orderEntity.ID,
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt.UnixNano(),
	}

	return order, res.Error
}

func (a Adapter) Save(order *domain.Order) error {
	var entityItems []OrderItem
	for _, orderItem := range order.OrderItems {
		entityItems = append(entityItems, OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
			OrderID:     uint(order.ID),
		})
	}
	orderM := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: entityItems,
	}
	res := a.db.Create(&orderM)
	if res.Error == nil {
		order.ID = uint(orderM.ID)
	}

	return res.Error
}
