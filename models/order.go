package models

import (
	"time"
)

type Order struct {
	ID           uint       `gorm:"primaryKey" db:"id" json:"id"`
	CustomerName string     `gorm:"not null" db:"customer_name" json:"customerName"`
	OrderedAt    *time.Time `gorm:"not null" db:"ordered_at" json:"orderedAt"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `db:"update_at" json:"update_at"`

	Items []OrderItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"items"`
}
