package ports

import (
	"context"

	"github.com/omareloui/gRPCMicroservicesInGo/services/payment/internal/application/core/domain"
)

type DBPort interface {
	Get(ctx context.Context, id uint) (domain.Payment, error)
	Save(ctx context.Context, payment *domain.Payment) error
}
