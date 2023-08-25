package usecase

import (
	"github.com/go-playground/validator/v10"
)

func (u UserUseCase) ValidateFields(payload interface{}) ([]map[string]string, error) {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		errorMessages := []map[string]string{}
		for _, fieldError := range err.(validator.ValidationErrors) {
			fieldName := fieldError.Field()
			errorMessage := ""
			switch fieldError.Tag() {
			case "required":
				errorMessage = fieldName + " is required"
			case "min":
				errorMessage = fieldName + " must have a minimum length of " + fieldError.Param()
			case "max":
				errorMessage = fieldName + " must have a maximum length of " + fieldError.Param()
			default:
				errorMessage = fieldError.Error()
			}

			errorMessages = append(errorMessages, map[string]string{
				fieldName: errorMessage,
			})
		}

		return errorMessages, err
	}
	return nil, nil
}
