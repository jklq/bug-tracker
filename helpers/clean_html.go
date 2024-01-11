package helpers

import (
	"github.com/microcosm-cc/bluemonday"
)

var s = bluemonday.StrictPolicy()

func CleanHTML(input string) string {
	return s.Sanitize(input)
}
