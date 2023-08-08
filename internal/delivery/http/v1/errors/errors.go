package errors

import (
	"errors"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
  Field string `json:"field"`
  Error string `json:"error"`
}

func AbortWithValidationErrors(ctx *gin.Context, err error) {

	var vd validator.ValidationErrors

	if errors.As(err, &vd) {
		payload := []ValidationError{}

		for _, v := range vd {
			payload = append(payload, ValidationError{strings.ToLower(v.Field()), v.Tag()})
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": payload,
			"type":   "VALIDATION_ERROR",
		})
		return
	}

	ctx.Status(http.StatusBadRequest)
}
