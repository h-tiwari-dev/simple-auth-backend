package utils

import (
	"sample-auth-backend/app/models"

	"github.com/go-playground/validator/v10"
)

// Custom validation function to check if either username or email is present
func usernameOrEmailValidation(fl validator.FieldLevel) bool {
	signIn, ok := fl.Parent().Interface().(models.SignIn)
	if !ok {
		return false
	}
	return signIn.Username != "" || signIn.Email != ""
}

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("username_or_email", usernameOrEmailValidation)

	return validate
}

func ValidatorError(err error) map[string]string {
	fields := map[string]string{}

	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}
	return fields
}
