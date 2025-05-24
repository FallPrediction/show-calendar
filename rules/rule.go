package rules

import (
	"show-calendar/initialize"
	"show-calendar/lang"
	en "show-calendar/lang/en"
	zhtw "show-calendar/lang/zh-tw"
	"strings"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"

	"github.com/go-playground/validator/v10"
)

type rule interface {
	validate() func(fl validator.FieldLevel) bool
}

func BindValidator(trans ut.Translator) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("exists", (&exists{}).validate())
		v.RegisterValidation("unique", (&unique{}).validate())
		registerTranslation(v, trans)
	}
}

// The params are splited by commas.
func getParams(fl *validator.FieldLevel) []string {
	return strings.Split((*fl).Param(), " ")
}

func registerTranslation(v *validator.Validate, trans ut.Translator) error {
	local := trans.Locale()
	var translations []lang.ValidationTrans
	switch local {
	case "zh_Hant_TW":
		translations = zhtw.ValidationTrans
	default:
		translations = en.ValidationTrans
	}

	var err error
	for _, t := range translations {
		if t.CustomTransFunc != nil && t.CustomRegisFunc != nil {
			err = v.RegisterTranslation(t.Tag, trans, t.CustomRegisFunc, t.CustomTransFunc)
		} else if t.CustomTransFunc != nil && t.CustomRegisFunc == nil {
			err = v.RegisterTranslation(t.Tag, trans, registrationFunc(t.Tag, t.Translation, t.Override), t.CustomTransFunc)
		} else if t.CustomTransFunc == nil && t.CustomRegisFunc != nil {
			err = v.RegisterTranslation(t.Tag, trans, t.CustomRegisFunc, translateFunc)
		} else {
			err = v.RegisterTranslation(t.Tag, trans, registrationFunc(t.Tag, t.Translation, t.Override), translateFunc)
		}
	}

	return err
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		if err = ut.Add(tag, translation, override); err != nil {
			return
		}
		return
	}

}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	logger := initialize.NewLogger()
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		logger.Error("警告: 翻譯欄位錯誤: %#v", fe)
		return fe.(error).Error()
	}

	return t
}
