package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Order struct {
	ID           uint       `gorm:"primaryKey" db:"id" json:"id"`
	CustomerName string     `gorm:"not null" db:"customer_name" json:"customerName" valid:"required"`
	OrderedAt    *time.Time `gorm:"not null" db:"ordered_at" json:"orderedAt" valid:"required"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `db:"update_at" json:"update_at"`

	Items []OrderItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"items"`
}

func (b *Order) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(b)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
