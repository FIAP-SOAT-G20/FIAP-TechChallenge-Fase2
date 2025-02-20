package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"

type PaymentExternal interface {
	CreatePayment(payment *entity.CreatePaymentIN) (*entity.CreatePaymentOUT, error)
	CreatePaymentMock(payment *entity.CreatePaymentIN) (*entity.CreatePaymentOUT, error)
}
