package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ErrorHandler struct {
	uni *ut.UniversalTranslator
}

func NewErrorHandler(uni *ut.UniversalTranslator) *ErrorHandler {
	return &ErrorHandler{
		uni: uni,
	}
}

func (h *ErrorHandler) HandleErrors(c *gin.Context) {
	c.Next()

	errorToPrint := c.Errors.ByType(gin.ErrorTypePublic).Last()
	if errorToPrint != nil {
		if errs, ok := errorToPrint.Err.(validator.ValidationErrors); ok {
			trans,_ := h.uni.GetTranslator("zh") // 这里也可以通过获取 HTTP Header 中的 Accept-Language 来获取用户的语言设置
			fmt.Println("aaa")
			c.JSON(http.StatusBadRequest, gin.H{
				"errors":  errs.Translate(trans),
			})
			return
		}

		// deal with other errors ...
	}
}