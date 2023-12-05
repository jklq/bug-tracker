package helpers

import (
	"fmt"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"

	enTranslations "github.com/go-playground/validator/v10/translations/en"

	ut "github.com/go-playground/universal-translator"
)

func TranslateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}

var Validate = validator.New()

var english = en.New()
var uni = ut.New(english, english)
var Translator, _ = uni.GetTranslator("en")

var _ = enTranslations.RegisterDefaultTranslations(Validate, Translator)
