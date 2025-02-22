package datasource

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type PaymentDataSource struct {
	db *gorm.DB
}

type PaymentKey string

func NewPaymentDataSource(db *gorm.DB) port.PaymentDataSource {
	return &PaymentDataSource{
		db: db,
	}
}

func (ds *PaymentDataSource) Create(ctx context.Context, payment *entity.Payment) (*entity.Payment, error) {
	if err := ds.db.WithContext(ctx).Create(payment).Error; err != nil {
		return nil, err
	}

	return payment, nil
}

func (ds *PaymentDataSource) GetPaymentByOrderIDAndStatus(ctx context.Context, status valueobject.PaymentStatus, orderID uint64) (*entity.Payment, error) {
	var payment entity.Payment

	if err := ds.db.WithContext(ctx).Where("order_id = ? AND status = ?", orderID, status).First(&payment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &payment, nil
		}
		return nil, err
	}

	return &payment, nil
}

func (ds *PaymentDataSource) UpdateStatus(status valueobject.PaymentStatus, externalPaymentID string) error {
	if err := ds.db.Model(&entity.Payment{}).Where("external_payment_id = ?", externalPaymentID).Update("status", status).Error; err != nil {
		return err
	}

	return nil
}
func (ds *PaymentDataSource) GetByExternalPaymentID(externalPaymentID string) (*entity.Payment, error) {
	var payment entity.Payment

	if err := ds.db.Where("external_payment_id = ?", externalPaymentID).First(&payment); errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	return &payment, nil
}
