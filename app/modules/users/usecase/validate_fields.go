package usecase

import (
	"github.com/go-playground/validator/v10"
)

func (u UserUseCase) ValidateFields(payload interface{}) (map[string]string, error) {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		errorMessages := make(map[string]string)
		for _, fieldError := range err.(validator.ValidationErrors) {
			fieldName := fieldError.Field()

			switch fieldError.Tag() {
			case "required":
				errorMessages[fieldName] = fieldName + " is required"
			case "min":
				errorMessages[fieldName] = fieldName + " must have a minimum length of " + fieldError.Param()
			case "max":
				errorMessages[fieldName] = fieldName + " must have a maximum length of " + fieldError.Param()
			default:
				errorMessages[fieldName] = fieldError.Error()
			}
		}
		return errorMessages, err
	}
	return nil, nil
}