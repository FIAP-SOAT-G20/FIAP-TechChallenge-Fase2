package entity

import (
	"time"

	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
)

type Payment struct {
	ID                uint64
	Status            valueobject.PaymentStatus
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
