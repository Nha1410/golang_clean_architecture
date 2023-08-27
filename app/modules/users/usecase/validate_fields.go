package usecase

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func (u UserUseCase) ValidateFields(payload interface{}) ([]map[string]string, error) {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		errorMessages := []map[string]string{}
		payloadType := reflect.TypeOf(payload).Elem()
		for _, fieldError := range err.(validator.ValidationErrors) {
			fieldName := fieldError.Field()
			field, _ := payloadType.FieldByName(fieldName)
			jsonTag := field.Tag.Get("json")
			errorMessage := ""
			switch fieldError.Tag() {
			case "required":
				errorMessage = jsonTag + " is required"
			case "min":
				errorMessage = jsonTag + " must have a minimum length of " + fieldError.Param()
			case "max":
				errorMessage = jsonTag + " must have a maximum length of " + fieldError.Param()
			default:
				errorMessage = fieldError.Error()
			}

			errorMessages = append(errorMessages, map[string]string{
				jsonTag: errorMessage,
			})
		}

		return errorMessages, err
	}
	return nil, nil
}
