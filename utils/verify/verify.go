package verify

import (
	"regexp"
)

func CheckMobile( mobile string) bool {
	reg := `^1\d{10}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(mobile)
}
