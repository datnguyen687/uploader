package delivery

import (
	"ecommerce/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      filter products
// @Description  filter products
// @Tags         products
// @Produce      json
// @Param        brand  query string false "brand"
// @Param        name  query string false "name"
// @Param        priceFrom  query int false "priceFrom"
// @Param        priceTo  query int false "priceTo"
// @Param		 orderBy query string false "orderBy" Enums(brand:asc, brand:desc, name:asc, name:desc, price:asc, price:desc)
// @Success      200  {object}  BasicResponse
// @Failure      400  {object}  BasicResponse
// @Failure      500  {object}  BasicResponse
// @Router       /v1/products [get]
func (d *httpDelivery) handleProducts(c *gin.Context) {
	var filter models.ProductFilter

	brand := c.Query("brand")
	if brand != "" {
		filter.Brand = &brand
	}

	name := c.Query("name")
	if name != "" {
		filter.Name = &name
	}

	priceFromStr := c.Query("priceFrom")
	if priceFromStr != "" {
		priceFrom, err := strconv.ParseFloat(priceFromStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, &BasicResponse{
				Code:    StatusCodeBadFilterProductsRequest,
				Message: StatusCodeBadFilterProductsRequest.Message(),
				Error:   err.Error(),
			})
		}
		filter.PriceFrom = &priceFrom
	}

	priceToStr := c.Query("priceTo")
	if priceToStr != "" {
		priceTo, err := strconv.ParseFloat(priceToStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, &BasicResponse{
				Code:    StatusCodeBadFilterProductsRequest,
				Message: StatusCodeBadFilterProductsRequest.Message(),
				Error:   err.Error(),
			})
		}
		filter.PriceTo = &priceTo
	}

	orderBy := c.Query("orderBy")

	products, err := d.uc.FilterProducts(c, filter, orderBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &BasicResponse{
			Code:    StatusCodeUnableToFilterProducts,
			Message: StatusCodeUnableToFilterProducts.Message(),
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &FilterProductsResponse{
		BasicResponse: BasicResponse{
			Code:    StatusCodeOk,
			Message: StatusCodeOk.Message(),
			Error:   "",
		},
		Data: products,
	})
}
