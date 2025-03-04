package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

var validate = validator.New()

func Init() {
	validate.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		slug := fl.Field().String()
		for _, char := range slug {
			if !(char == '_' || (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
				return false
			}
		}
		return true
	})
}

func DecodeJSONBody(r *http.Request, input any) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(input); err != nil {
		return err
	}
	if err := validate.Struct(input); err != nil {
		return err
	}
	return nil
}
