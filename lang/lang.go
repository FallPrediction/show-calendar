package lang

import "github.com/go-playground/validator/v10"

type ValidationTrans struct {
	Tag             string
	Translation     string
	Override        bool
	CustomRegisFunc validator.RegisterTranslationsFunc
	CustomTransFunc validator.TranslationFunc
}
