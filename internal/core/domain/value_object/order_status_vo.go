package valueobject

import (
	"strings"
)

type OrderStatus string

const (
	UNDEFINDED OrderStatus = "UNDEFINDED"
	OPEN       OrderStatus = "OPEN"
	CANCELLED  OrderStatus = "CANCELLED"
	PENDING    OrderStatus = "PENDING"
	RECEIVED   OrderStatus = "RECEIVED"
	PREPARING  OrderStatus = "PREPARING"
	READY      OrderStatus = "READY"
	COMPLETED  OrderStatus = "COMPLETED"
)

func IsValidOrderStatus(status string) bool {
	return ToOrderStatus(status) != UNDEFINDED
}

// String returns the string representation of the OrderStatus
func (o OrderStatus) String() string {
	return strings.ToUpper(string(o))
}

// ToOrderStatus converts a string to an OrderStatus
func ToOrderStatus(status string) OrderStatus {
	switch strings.ToUpper(status) {
	case "OPEN":
		return OPEN
	case "CANCELLED":
		return CANCELLED
	case "PENDING":
		return PENDING
	case "RECEIVED":
		return RECEIVED
	case "PREPARING":
		return PREPARING
	case "READY":
		return READY
	case "COMPLETED":
		return COMPLETED
	default:
		return UNDEFINDED
	}
}

// OrderStatusTransitions defines the allowed transitions between OrderStatuses
var OrderStatusTransitions = map[OrderStatus][]OrderStatus{
	OPEN:      {CANCELLED, PENDING},
	CANCELLED: {},
	PENDING:   {OPEN, RECEIVED},
	RECEIVED:  {PREPARING},
	PREPARING: {READY},
	READY:     {COMPLETED},
	COMPLETED: {},
}

// StatusCanTransitionTo returns true if the transition from oldStatus to newStatus is allowed
func StatusCanTransitionTo(oldStatus, newStatus OrderStatus) bool {
	allowedStatuses := OrderStatusTransitions[oldStatus]
	for _, allowedStatus := range allowedStatuses {
		if newStatus == allowedStatus {
			return true
		}
	}
	return false
}

// StatusTransitionNeedsStaffID returns true if the new status requires a staff ID
func StatusTransitionNeedsStaffID(newStatus OrderStatus) bool {
	switch newStatus {
	case OPEN:
		return false
	case CANCELLED:
		return false
	case PENDING:
		return false
	case RECEIVED:
		return false
	case PREPARING:
		return true
	case READY:
		return true
	case COMPLETED:
		return true
	default:
		return false
	}
}
