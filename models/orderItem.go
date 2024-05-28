package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID          uint   `gorm:"primaryKey" db:"id" json:"id"`
	Name        string `gorm:"not null" db:"name" json:"name" form:"name" valid:"required~Name of product is required"`
	Description string `gorm:"not null" db:"description" json:"description" form:"description" valid:"required~Description of product is required"`
	Quantity    int    `gorm:"not null" db:"quantity" json:"quantity" form:"quantity" valid:"required~Quantity of product is required, numeric~Invalid quantity format"`
	OrderID     uint   `db:"order_id"`
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
