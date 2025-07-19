package rules

import (
	"fmt"
	"souflair/config"
	"souflair/initialize"

	"github.com/go-playground/validator/v10"
)

var _ rule = new(unique)

type unique struct{}

func (r *unique) validate() func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		param := getParams(&fl)
		value := fl.Field().String()
		pg := config.NewPg()
		db := initialize.NewDB(pg)
		var count int64
		if db.Table(param[0]).Where(fmt.Sprintf("%s = ?", param[1]), value).Count(&count); count == 0 {
			return true
		}
		return false
	}
}
