package error

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/search-n-cache/search-n-cache-service/response"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)


func HandleError(err error, ctx *gin.Context) {
	errorId := uuid.NewString()
	log.Errorf("Error [id = %s] occured while processing request: [msg = %s]", errorId, err)
	httpStatus := http.StatusInternalServerError
	// specific error check
	if errors.Is(err, gorm.ErrRecordNotFound) {
		httpStatus = http.StatusNotFound
	}
	// error type check
	if errors.As(err, &validator.ValidationErrors{}) {
		httpStatus = http.StatusBadRequest
	}
	res := response.ErrorResponse{
		ErrorCode:    http.StatusText(httpStatus),
		ErrorMessage: err.Error(),
		ErrorId: errorId,
	}
	ctx.JSON(httpStatus, res)
}