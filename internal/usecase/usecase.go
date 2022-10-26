package usecase

import (
	"context"
	"ecommerce/internal/models"
	"ecommerce/internal/repository"
)

func NewUsecase(repo repository.IRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}

type IUsecase interface {
	FilterProducts(ctx context.Context, filter models.ProductFilter, orderBy string) ([]models.Product, error)
}

type usecase struct {
	repo repository.IRepository
}
