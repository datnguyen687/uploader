package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      user batch
// @Description  user batch
// @Tags         user
// @Accept 		 multipart/form-data
// @Produce      json
// @Param		 file formData file true "file"
// @Success      200  {object}  BasicResponse
// @Failure      400  {object}  BasicResponse
// @Failure      500  {object}  BasicResponse
// @Router       /v1/user/batch [post]
func (d *httpDelivery) batchUser(c *gin.Context) {
	multiPartHeaders, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, &BasicResponse{
			Code:    StatusCodeBadBatchUserRequest,
			Message: StatusCodeBadBatchUserRequest.Message(),
			Error:   err.Error(),
		})
		return
	}

	if multiPartHeaders.Size > maxFileSizeInBytes {
		c.JSON(http.StatusBadRequest, &BasicResponse{
			Code:    StatusCodeBadBatchUserRequest,
			Message: StatusCodeBadBatchUserRequest.Message(),
			Error:   "file is too big",
		})
		return
	}

	file, err := multiPartHeaders.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, &BasicResponse{
			Code:    StatusCodeBadBatchUserRequest,
			Message: StatusCodeBadBatchUserRequest.Message(),
			Error:   err.Error(),
		})
		return
	}

	if err := d.uc.Upload(c, multiPartHeaders.Filename, file); err != nil {
		c.JSON(http.StatusInternalServerError, &BasicResponse{
			Code:    StatusCodeUnableToBatchUser,
			Message: StatusCodeUnableToBatchUser.Message(),
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &BasicResponse{
		Code:    StatusCodeOk,
		Message: StatusCodeOk.Message(),
		Error:   "",
	})
}
