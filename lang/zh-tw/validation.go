package zhtw

import (
	"souflair/lang"
)

var ValidationTrans = []lang.ValidationTrans{
	{
		Tag:         "unique",
		Translation: "{0}已存在",
		Override:    false,
	},
	{
		Tag:         "exists",
		Translation: "{0}不存在",
		Override:    false,
	},
}
