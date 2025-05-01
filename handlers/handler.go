package handlers

import (
	"errors"
	ut "github.com/go-playground/universal-translator"
	"math"
	"net/http"
	"show-calendar/initialize"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Handler struct {
	translator ut.Translator
}

func (handler *Handler) handleErrorAndReturn(c *gin.Context, err error, onSuccess func()) {
	if err != nil {
		handler.handleError(c, err)
		return
	}
	onSuccess()
}

func (handler *Handler) handleError(c *gin.Context, err error) {
	logger := initialize.NewLogger()
	switch {
	case errors.As(err, &(validator.ValidationErrors{})):
		handler.sendResponse(c, http.StatusUnprocessableEntity, err.(validator.ValidationErrors).Translate(handler.translator), nil)
	case errors.Is(err, gorm.ErrRecordNotFound):
		handler.sendResponse(c, http.StatusNotFound, "伺服器找不到請求的資源", nil)
	default:
		logger.Error("system error ", err)
		handler.sendResponse(c, http.StatusUnprocessableEntity, "系統錯誤", nil)
	}
}

func (handler *Handler) sendResponse(c *gin.Context, statusCode int, message interface{}, data interface{}) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}

func (handler *Handler) sendResponseWithPagination(c *gin.Context, statusCode int, message interface{}, data interface{}, currentPage int, perPage int, total int) {
	lastPage := int(math.Ceil(float64(total) / float64(perPage)))
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
		"pagination": gin.H{
			"currentPage": currentPage,
			"lastPage":    lastPage,
			"total":       total,
			"perPage":     perPage,
		},
	})
}

func NewBaseHandler(translator ut.Translator) Handler {
	return Handler{translator: translator}
}
