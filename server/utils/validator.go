package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(strct any) string {
	msg := ""
	err := validate.Struct(strct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			msg += fmt.Sprintf("%s %s %s ", err.StructNamespace(), err.Tag(), err.Param())
		}
	}
	return msg
}

func ValidateStructInRequest(strct any, c *fiber.Ctx) error {
	errs := ValidateStruct(strct)

	if len(errs) != 0 {
		c.JSON(fiber.Map{
			"message": errs,
		})
		return errors.New(errs)
	}

	return nil
}
