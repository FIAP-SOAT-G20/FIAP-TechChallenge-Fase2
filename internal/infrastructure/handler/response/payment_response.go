package response

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"

type CreatePaymentResponse struct {
	InStoreOrderID string `json:"in_store_order_id"`
	QrData         string `json:"qr_data"`
}

func ToCreatePaymentOUTDomain(payment *CreatePaymentResponse) *entity.CreatePaymentOutput {
	return &entity.CreatePaymentOutput{
		InStoreOrderID: payment.InStoreOrderID,
		QrData:         payment.QrData,
	}
}
