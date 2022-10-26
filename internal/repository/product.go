package repository

import (
	"context"
	"ecommerce/internal/models"
	"fmt"
)

func (repo *repository) FilterProducts(ctx context.Context, filter models.ProductFilter, orderBy *models.ProductOrderBy) ([]models.Product, error) {
	db := repo.db

	if filter.Brand != nil {
		db = db.Where("brand = ?", *filter.Brand)
	}

	if filter.Name != nil {
		db = db.Where("name = ?", *filter.Name)
	}

	if filter.PriceFrom != nil {
		db = db.Where("price >= ?", *filter.PriceFrom)
	}

	if filter.PriceTo != nil {
		db = db.Where("price <= ?", *filter.PriceTo)
	}

	if orderBy != nil {
		db = db.Order(fmt.Sprintf("%s %s", orderBy.Column, orderBy.Order))
	}

	var products []models.Product
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
