package delivery_test

import (
	"ecommerce/internal/delivery"
	"ecommerce/internal/models"
	"ecommerce/mocks"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPingOk(t *testing.T) {
	engine := gin.Default()

	uc := mocks.IUsecase{}

	d := delivery.NewHttpDelivery(&uc)
	d.Init(engine, true, false)

	expectedResp := delivery.BasicResponse{
		Code:    delivery.StatusCodeOk,
		Message: delivery.StatusCodeOk.Message(),
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/v1/ping", nil)
	assert.NoError(t, err)

	engine.ServeHTTP(w, req)

	data, err := ioutil.ReadAll(w.Body)
	assert.Nil(t, err)

	var resp delivery.BasicResponse
	json.Unmarshal(data, &resp)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedResp, resp)
}

func TestHandleProductsOk(t *testing.T) {
	engine := gin.Default()

	name := "fake name"
	brand := "fake brand"
	priceFrom := float64(1)
	priceTo := float64(1000)

	orderBy := "name:desc"
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

	uc := mocks.IUsecase{}
	uc.On("FilterProducts", mock.Anything, mock.AnythingOfType("models.ProductFilter"), orderBy).Return(expectedProducts, nil)

	d := delivery.NewHttpDelivery(&uc)
	d.Init(engine, true, false)

	expectedResp := delivery.FilterProductsResponse{
		BasicResponse: delivery.BasicResponse{
			Code:    delivery.StatusCodeOk,
			Message: delivery.StatusCodeOk.Message(),
		},
		Data: expectedProducts,
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/v1/products", nil)
	assert.NoError(t, err)

	query := req.URL.Query()
	query.Add("name", name)
	query.Add("brand", brand)
	query.Add("priceFrom", fmt.Sprintf("%f", priceFrom))
	query.Add("priceTo", fmt.Sprintf("%f", priceTo))
	query.Add("orderBy", orderBy)
	req.URL.RawQuery = query.Encode()

	assert.NoError(t, err)

	engine.ServeHTTP(w, req)

	data, err := ioutil.ReadAll(w.Body)
	assert.Nil(t, err)

	var resp delivery.FilterProductsResponse
	json.Unmarshal(data, &resp)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedResp, resp)

}

func TestHandleProductsErrInvalidRequest(t *testing.T) {
	engine := gin.Default()

	name := "fake name"
	brand := "fake brand"
	priceFrom := "abcd"
	priceTo := "abcd"

	orderBy := "name:desc"
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

	uc := mocks.IUsecase{}
	uc.On("FilterProducts", mock.Anything, mock.AnythingOfType("models.ProductFilter"), orderBy).Return(expectedProducts, nil)

	d := delivery.NewHttpDelivery(&uc)
	d.Init(engine, true, false)

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/v1/products", nil)
	assert.NoError(t, err)

	query := req.URL.Query()
	query.Add("name", name)
	query.Add("brand", brand)
	query.Add("priceFrom", priceFrom)
	query.Add("priceTo", priceTo)
	query.Add("orderBy", orderBy)
	req.URL.RawQuery = query.Encode()

	assert.NoError(t, err)

	engine.ServeHTTP(w, req)

	data, err := ioutil.ReadAll(w.Body)
	assert.Nil(t, err)

	var resp delivery.FilterProductsResponse
	json.Unmarshal(data, &resp)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandleProductsErrUnableToFilter(t *testing.T) {
	engine := gin.Default()

	name := "fake name"
	brand := "fake brand"

	orderBy := "name:desc"

	uc := mocks.IUsecase{}
	uc.On("FilterProducts", mock.Anything, mock.AnythingOfType("models.ProductFilter"), orderBy).Return(nil, errors.New("unable to filter products"))

	d := delivery.NewHttpDelivery(&uc)
	d.Init(engine, true, false)

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/v1/products", nil)
	assert.NoError(t, err)

	query := req.URL.Query()
	query.Add("name", name)
	query.Add("brand", brand)
	query.Add("orderBy", orderBy)
	req.URL.RawQuery = query.Encode()

	assert.NoError(t, err)

	engine.ServeHTTP(w, req)

	data, err := ioutil.ReadAll(w.Body)
	assert.Nil(t, err)

	var resp delivery.FilterProductsResponse
	json.Unmarshal(data, &resp)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
