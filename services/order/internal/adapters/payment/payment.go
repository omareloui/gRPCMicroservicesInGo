package payment

import (
	"context"

	"github.com/omareloui/gRPCMicroservicesInGo/proto/go/payment"
	"github.com/omareloui/gRPCMicroservicesInGo/services/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment payment.PaymentServiceClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := payment.NewPaymentServiceClient(conn)
	return &Adapter{payment: client}, err
}

func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.payment.CreatePayment(context.Background(), &payment.CreatePaymentRequest{
		UserId:  int64(order.CustomerID),
		OrderId: int64(order.ID),
		Price:   order.TotalPrice(),
	})
	return err
}
