package usecase

import (
	"context"
	"ecommerce/internal/models"
	"errors"
	"strings"
)

var (
	columnMapping = map[string]string{
		"brand": "brand",
		"name":  "name",
		"price": "price",
	}
	orderMapping = map[string]string{
		"desc": "DESC",
		"asc":  "ASC",
	}
)

func (uc *usecase) FilterProducts(ctx context.Context, filter models.ProductFilter, orderBy string) ([]models.Product, error) {
	var productOrderBy *models.ProductOrderBy
	if orderBy != "" {
		tokens := strings.Split(orderBy, ":")
		if len(tokens) < 2 {
			return nil, errors.New("invalid orderBy")
		}

		columnName, ok := columnMapping[tokens[0]]
		if !ok {
			return nil, errors.New("invalid orderBy")
		}

		order, ok := orderMapping[tokens[1]]
		if !ok {
			return nil, errors.New("invalid orderBy")
		}

		productOrderBy = &models.ProductOrderBy{
			Column: columnName,
			Order:  order,
		}
	}
	return uc.repo.FilterProducts(ctx, filter, productOrderBy)
}
