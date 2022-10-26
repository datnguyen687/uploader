package repository

import (
	"context"
	"ecommerce/internal/models"

	"gorm.io/gorm"
)

type IRepository interface {
	FilterProducts(ctx context.Context, filter models.ProductFilter, orderBy *models.ProductOrderBy) ([]models.Product, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *gorm.DB
}

func (repo *repository) AutoMigrate() error {
	return repo.db.AutoMigrate(
		&models.Product{},
	)
}
