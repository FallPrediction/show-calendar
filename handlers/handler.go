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

func (h *Handler) handleErrorAndReturn(c *gin.Context, err error, onSuccess func()) {
	if err != nil {
		h.handleError(c, err)
		return
	}
	onSuccess()
}

func (h *Handler) handleError(c *gin.Context, err error) {
	logger := initialize.NewLogger()
	switch {
	case errors.As(err, &(validator.ValidationErrors{})):
		h.sendResponse(c, http.StatusUnprocessableEntity, err.(validator.ValidationErrors).Translate(h.translator), nil)
	case errors.Is(err, gorm.ErrRecordNotFound):
		h.sendResponse(c, http.StatusNotFound, "伺服器找不到請求的資源", nil)
	default:
		logger.Error("system error ", err)
		h.sendResponse(c, http.StatusUnprocessableEntity, "系統錯誤", nil)
	}
}

func (h *Handler) sendResponse(c *gin.Context, statusCode int, message interface{}, data interface{}) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}

func (h *Handler) sendResponseWithPagination(c *gin.Context, statusCode int, message interface{}, data interface{}, currentPage int, perPage int, total int) {
	lastPage := int(math.Ceil(float64(total) / float64(perPage)))
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
		"pagination": gin.H{
			"current_page": currentPage,
			"last_page":    lastPage,
			"total":        total,
			"per_page":     perPage,
		},
	})
}

func NewBaseHandler(translator ut.Translator) Handler {
	return Handler{translator: translator}
}
