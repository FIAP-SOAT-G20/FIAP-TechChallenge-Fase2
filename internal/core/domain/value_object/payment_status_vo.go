package valueobject

import "strings"

type PaymentStatus string

const (
	PROCESSING   PaymentStatus = "PROCESSING"
	CONFIRMED    PaymentStatus = "CONFIRMED"
	FAILED       PaymentStatus = "FAILED"
	CANCELED     PaymentStatus = "CANCELED"
	UNDEFINDED_P PaymentStatus = ""
)

func IsValidPaymentStatus(status string) bool {
	return ToPaymentStatus(status) != UNDEFINDED_P
}

// String returns the string representation of the PaymentStatus
func (o PaymentStatus) String() string {
	return strings.ToUpper(string(o))
}

// ToPaymentStatus converts a string to an PaymentStatus
func ToPaymentStatus(status string) PaymentStatus {
	switch strings.ToUpper(status) {
	case "PROCESSING":
		return PROCESSING
	case "CONFIRMED":
		return CONFIRMED
	case "FAILED":
		return FAILED
	case "CANCELED":
		return CANCELED
	default:
		return UNDEFINDED_P
	}
}
