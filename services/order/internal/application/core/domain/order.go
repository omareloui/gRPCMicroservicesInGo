package domain

import "time"

type OrderItem struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    uint32  `json:"quantity"`
}

type Order struct {
	ID         uint        `json:"id"`
	CustomerID uint        `json:"customer_id"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  int64       `json:"created_at"`
}

func NewOrder(customerId uint, orderItems []OrderItem) Order {
	return Order{
		CreatedAt:  time.Now().Unix(),
		Status:     "pending",
		CustomerID: customerId,
		OrderItems: orderItems,
	}
}

func (o *Order) TotalPrice() float32 {
	sum := float32(0)
	for _, item := range o.OrderItems {
		sum += item.UnitPrice * float32(item.Quantity)
	}

	return sum
}
