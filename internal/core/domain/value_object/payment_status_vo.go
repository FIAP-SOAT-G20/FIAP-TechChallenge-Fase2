package valueobject

type PaymentStatus string

const (
	PROCESSING PaymentStatus = "PROCESSING"
	CONFIRMED  PaymentStatus = "CONFIRMED"
	FAILED     PaymentStatus = "FAILED"
	CANCELED   PaymentStatus = "CANCELED"
)
