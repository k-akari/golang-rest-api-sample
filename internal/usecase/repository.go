//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock
package usecase

import (
	"context"
	"time"

	"github.com/k-akari/golang-rest-api-sample/internal/domain"
)

type (
	companyRepository interface {
		Create(ctx context.Context, company *domain.Company) (domain.CompanyID, error)
		GetByID(ctx context.Context, companyID domain.CompanyID) (*domain.Company, error)
	}
	clientRepository interface {
		Create(ctx context.Context, client *domain.Client) (domain.ClientID, error)
		GetByID(ctx context.Context, clientID domain.ClientID) (*domain.Client, error)
	}
	invoiceRepository interface {
		Create(ctx context.Context, invoice *domain.Invoice) (domain.InvoiceID, error)
		ListByPaymentDueDateBetween(ctx context.Context, coid domain.CompanyID, from, to *time.Time) ([]*domain.Invoice, error)
	}
)
