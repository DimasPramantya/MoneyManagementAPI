package validation

import (
	"fmt"
	
	. "github.com/dimas-pramantya/money-management/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (v *Validator) ValidateRequest(ctx *gin.Context, req any) error {
	if err := ctx.ShouldBindJSON(req); err != nil {
		err := BadRequestError("Invalid request", []string{"Invalid Content-Type, expected application/json"})
		ctx.Error(err)
	}

	if errs := v.Validate(req); errs != nil {
		err := BadRequestError("Invalid request", errs)
		ctx.Error(err)
	}

	return nil
}

func (v *Validator) ValidateQuery(ctx *gin.Context, req any) error {
	if err := ctx.ShouldBindQuery(req); err != nil {
		err := BadRequestError("Invalid request", []string{"Invalid Content-Type, expected application/json"})
		ctx.Error(err)
	}
	if errs := v.Validate(req); errs != nil {
		err := BadRequestError("Invalid request", errs)
		ctx.Error(err)
	}
	return nil
}

func (v *Validator) Validate(i any) []string {
	err := v.validator.Struct(i)
	if err == nil {
		return nil
	}

	var data []string
	for _, e := range err.(validator.ValidationErrors) {
		data = append(data, fmt.Sprintf("%s is %s", e.Field(), e.Tag()))
	}
	return data
}
