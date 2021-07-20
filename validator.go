package main

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"log"
)

type Validator struct {
	Translator ut.Translator
	Validate   *validator.Validate
}

func NewValidator() (*Validator, error) {
	validate := validator.New()

	enLocale := en.New()
	u := ut.New(enLocale, enLocale)

	trans, _ := u.GetTranslator("en")

	err := enTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return nil, err
	}

	v := &Validator{
		Translator: trans,
		Validate:   validate,
	}

	err = v.RegisterTranslations()
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (v *Validator) Map(err error) map[string]interface{} {
	m := map[string]interface{}{}

	validationErrors := err.(validator.ValidationErrors)

	for _, e := range validationErrors {
		m[e.Field()] = e.Translate(v.Translator)
	}

	return m
}

func (v *Validator) RegisterTranslations() error {
	// required
	err := v.Validate.RegisterTranslation(
		"required",
		v.Translator,
		func(u ut.Translator) error {
			return u.Add("required", "This field is required.", true)
		},
		func(u ut.Translator, fe validator.FieldError) string {
			t, err := u.T("required", fe.Field())
			if err != nil {
				log.Panic(err)
			}

			return t
		})
	if err != nil {
		return err
	}

	// url
	err = v.Validate.RegisterTranslation(
		"url",
		v.Translator,
		func(u ut.Translator) error {
			return u.Add("url", "This field should be a valid URL.", true)
		},
		func(u ut.Translator, fe validator.FieldError) string {
			t, err := u.T("url", fe.Field())
			if err != nil {
				log.Panic(err)
			}

			return t
		})
	if err != nil {
		return err
	}

	return nil
}
