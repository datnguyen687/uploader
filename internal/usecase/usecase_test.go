package usecase_test

import (
	"context"
	"ecommerce/internal/models"
	"ecommerce/internal/usecase"
	"ecommerce/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFilterProductsOk(t *testing.T) {
	ctx := context.Background()
	name := "fake name"
	brand := "fake brand"
	priceFrom := float64(1)
	priceTo := float64(1000)

	filter := models.ProductFilter{
		Name:      &name,
		Brand:     &brand,
		PriceFrom: &priceFrom,
		PriceTo:   &priceTo,
	}

	orderByStr := "name:desc"
	orderBy := models.ProductOrderBy{
		Column: "name",
		Order:  "DESC",
	}

	now := time.Now().UTC()

	expectedProducts := []models.Product{
		{
			ID:        1,
			Name:      "fake name",
			Brand:     "fake brand",
			Price:     100,
			CreatedAt: now,
		},
	}

	mockRepo := mocks.IRepository{}
	mockRepo.On("FilterProducts", mock.Anything, filter, &orderBy).Return(expectedProducts, nil)

	uc := usecase.NewUsecase(&mockRepo)

	products, err := uc.FilterProducts(ctx, filter, orderByStr)
	assert.NoError(t, err)

	assert.NotNil(t, products)
	assert.Equal(t, len(expectedProducts), len(products))
}

func TestFilterProductsErrUnableToFilter(t *testing.T) {
	ctx := context.Background()
	name := "fake name"
	brand := "fake brand"
	priceFrom := float64(1)
	priceTo := float64(1000)

	filter := models.ProductFilter{
		Name:      &name,
		Brand:     &brand,
		PriceFrom: &priceFrom,
		PriceTo:   &priceTo,
	}

	orderByStr := "name:desc"
	orderBy := models.ProductOrderBy{
		Column: "name",
		Order:  "DESC",
	}

	expectedErr := errors.New("unable to find")

	mockRepo := mocks.IRepository{}
	mockRepo.On("FilterProducts", mock.Anything, filter, &orderBy).Return(nil, expectedErr)

	uc := usecase.NewUsecase(&mockRepo)

	products, err := uc.FilterProducts(ctx, filter, orderByStr)
	assert.Equal(t, expectedErr, err)

	assert.Nil(t, products)
}

func TestFilterProductsErrInvalidOrderBy(t *testing.T) {
	ctx := context.Background()
	name := "fake name"
	brand := "fake brand"
	priceFrom := float64(1)
	priceTo := float64(1000)

	filter := models.ProductFilter{
		Name:      &name,
		Brand:     &brand,
		PriceFrom: &priceFrom,
		PriceTo:   &priceTo,
	}

	orderByStr := "abcd:efg"

	expectedErr := errors.New("invalid orderBy")

	mockRepo := mocks.IRepository{}

	uc := usecase.NewUsecase(&mockRepo)

	products, err := uc.FilterProducts(ctx, filter, orderByStr)
	assert.Equal(t, expectedErr, err)

	assert.Nil(t, products)
}
