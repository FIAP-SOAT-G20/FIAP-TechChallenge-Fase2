package entity

import "time"

type PaymentStatus string

const (
	PROCESSING PaymentStatus = "PROCESSING"
	CONFIRMED  PaymentStatus = "CONFIRMED"
	FAILED     PaymentStatus = "FAILED"
	CANCELED   PaymentStatus = "CANCELED"
)

type Payment struct {
	ID                uint64
	Status            PaymentStatus
	ExternalPaymentID string
	QrData            string
	OrderID           uint64
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type CreatePaymentIN struct {
	ExternalReference string
	TotalAmount       float32
	Items             []ItemsIN
	Title             string
	Description       string
	NotificationUrl   string
}

type ItemsIN struct {
	Category    string
	Title       string
	Description string
	UnitPrice   float32
	Quantity    uint64
	UnitMeasure string
	TotalAmount float32
}

type CreatePaymentOUT struct {
	InStoreOrderID string
	QrData         string
}

type UpdatePaymentIN struct {
	Resource string
	Topic    string
}

type GetPaymentOUT struct {
	ExternalReference string
}

type CreatePaymentRequest struct {
	ExternalReference string         `json:"external_reference"`
	TotalAmount       float32        `json:"total_amount"`
	Items             []ItemsRequest `json:"items"`
	Title             string         `json:"title"`
	Description       string         `json:"description"`
	NotificationURL   string         `json:"notification_url"`
}

type ItemsRequest struct {
	Category    string  `json:"category"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    uint64  `json:"quantity"`
	UnitMeasure string  `json:"unit_measure"`
	TotalAmount float32 `json:"total_amount"`
}

func NewPaymentRequest(payment *CreatePaymentIN) *CreatePaymentRequest {
	if payment == nil {
		return nil
	}

	items := make([]ItemsRequest, 0)
	for _, item := range payment.Items {
		items = append(items, ItemsRequest{
			Title:       item.Title,
			Description: item.Description,
			UnitPrice:   item.UnitPrice,
			Category:    item.Category,
			UnitMeasure: item.UnitMeasure,
			Quantity:    item.Quantity,
			TotalAmount: item.TotalAmount,
		})
	}

	return &CreatePaymentRequest{
		ExternalReference: payment.ExternalReference,
		TotalAmount:       payment.TotalAmount,
		Items:             items,
		Title:             payment.Title,
		Description:       payment.Description,
		NotificationURL:   payment.NotificationUrl,
	}
}

func ToCreatePaymentOUTDomain(payment *CreatePaymentResponse) *CreatePaymentOUT {
	return &CreatePaymentOUT{
		InStoreOrderID: payment.InStoreOrderID,
		QrData:         payment.QrData,
	}
}

type CreatePaymentResponse struct {
	InStoreOrderID string `json:"in_store_order_id"`
	QrData         string `json:"qr_data"`
}

type PaymentResponse struct {
	ID                uint64        `json:"id"`
	Status            PaymentStatus `json:"status"`
	OrderID           uint64        `json:"order_id"`
	ExternalPaymentID string        `json:"external_payment_id"`
	QrData            string        `json:"qr_data"`
}
