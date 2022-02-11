package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(100); not null"`
	Observations *string `gorm:"type:varchar(100);"`
	Price        float64 `gorm:"type:float; not null"`
}
type InvoiceHeader struct {
	gorm.Model
	Client       string `gorm:"type:varchar(200); not null"`
	InvoiceItems []InvoiceItem
}

type InvoiceItem struct {
	gorm.Model
	InvoiceHeaderID uint
	ProductID       uint
}
