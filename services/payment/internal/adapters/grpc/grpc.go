package grpc

import (
	"context"
	"fmt"

	"github.com/omareloui/gRPCMicroservicesInGo/proto/go/payment"

	"github.com/omareloui/gRPCMicroservicesInGo/services/payment/internal/application/core/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	newPayment := domain.NewPayment(uint(request.UserId), uint(request.OrderId), request.Price)
	result, err := a.api.Charge(ctx, newPayment)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &payment.CreatePaymentResponse{PaymentId: int64(result.ID)}, nil
}
