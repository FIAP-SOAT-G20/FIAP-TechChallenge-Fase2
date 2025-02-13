package entity

import (
	"time"
)

type Role string

const (
	COOK      Role = "COOK"
	ATTENDANT Role = "ATTENDANT"
	MANAGER   Role = "MANAGER"
)

type Staff struct {
	ID        uint64
	Name      string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewStaff(name string, role Role) *Staff {
	staff := &Staff{
		Name:      name,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return staff
}

func (p *Staff) Update(name string, role Role) {
	p.Name = name
	p.Role = role
	p.UpdatedAt = time.Now()
}
