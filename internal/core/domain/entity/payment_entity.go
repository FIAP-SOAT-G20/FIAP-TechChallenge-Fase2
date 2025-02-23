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

type CreatePaymentInput struct {
	ExternalReference string
	TotalAmount       float32
	Items             []ItemsInput
	Title             string
	Description       string
	NotificationUrl   string
}

type ItemsInput struct {
	Category    string
	Title       string
	Description string
	UnitPrice   float32
	Quantity    uint64
	UnitMeasure string
	TotalAmount float32
}

type CreatePaymentOutput struct {
	InStoreOrderID string
	QrData         string
}

type UpdatePaymentInput struct {
	Resource string
	Topic    string
}

type GetPaymentOutput struct {
	ExternalReference string
}
