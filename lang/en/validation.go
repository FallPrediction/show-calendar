package en

import (
	"souflair/lang"
)

var ValidationTrans = []lang.ValidationTrans{
	{
		Tag:         "unique",
		Translation: "The {0} has already been taken.",
		Override:    false,
	},
	{
		Tag:         "exists",
		Translation: "The selected {0} is invalid.",
		Override:    false,
	},
}
