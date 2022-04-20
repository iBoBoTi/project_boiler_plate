package validation

import (
	"errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	goValidator "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
)

type goPlayground struct {
	validator *goValidator.Validate
	translate ut.Translator
	err       error
	msg       []string
}

func NewGoPlayground() (ports.Validator, error) {
	var (
		language         = en.New()
		uni              = ut.New(language, language)
		translate, found = uni.GetTranslator("en")
	)

	if !found {
		return nil, errors.New("translator not found")
	}

	v := goValidator.New()
	if err := enTranslations.RegisterDefaultTranslations(v, translate); err != nil {
		return nil, errors.New("translator not found")
	}

	return &goPlayground{validator: v, translate: translate}, nil
}

func (g *goPlayground) Validate(i interface{}) error {
	if len(g.msg) > 0 {
		g.msg = nil
	}

	g.err = g.validator.Struct(i)
	if g.err != nil {
		return g.err
	}

	return nil
}

func (g *goPlayground) Messages() []string {
	if g.err != nil {
		for _, err := range g.err.(goValidator.ValidationErrors) {
			g.msg = append(g.msg, err.Translate(g.translate))
		}
	}

	return g.msg
}
