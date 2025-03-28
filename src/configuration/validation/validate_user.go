package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_transalation "github.com/go-playground/validator/v10/translations/en"
	resterr "github.com/nfdeveloper/crud_with_authentication/src/configuration/rest_err"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {

		en := en.New()
		unt := ut.New(en, en)

		transl, _ = unt.GetTranslator("en")
		en_transalation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *resterr.RestErr {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return resterr.NewBadRequestError("Invalid field type.")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []resterr.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := resterr.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return resterr.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return resterr.NewBadRequestError("Error trying to convert fields")
	}
}
