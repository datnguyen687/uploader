package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type ProductFilter struct {
	Name      *string
	Brand     *string
	PriceFrom *float64
	PriceTo   *float64
}

type ProductOrderBy struct {
	Column string
	Order  string // DESC/ASC
}

type Product struct {
	ID    int     `gorm:"column:id;type:serial;primaryKey"`
	Name  string  `gorm:"column:name;type:text;NOT NULL"`
	Brand string  `gorm:"column:brand;type:text;NOT NULL"`
	Price float64 `gorm:"column:price;type:numeric;default: 0"`

	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp;default: now()"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp"`
}

func (Product) Tablename() string {
	return "products"
}

func (*Product) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.New())
}
