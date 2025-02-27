package entity

import "time"

type Category struct {
	ID        uint64
	Name      string
	CreatedAt time.Time
}
