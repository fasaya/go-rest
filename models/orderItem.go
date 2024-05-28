package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID          uint   `gorm:"primaryKey" db:"id" json:"id"`
	Name        string `gorm:"not null" db:"name" json:"name" valid:"required"`
	Description string `gorm:"not null" db:"description" json:"description" valid:"required"`
	Quantity    int    `gorm:"not null" db:"quantity" json:"quantity" valid:"required,numeric"`
	OrderID     uint   `gorm:"not null" db:"order_id"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

func (b *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(b)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
