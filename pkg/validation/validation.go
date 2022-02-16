package validation

import (
	"regexp"
)

func FieldLen(min, max int, name string) bool {
	if len(name) < min || len(name) > max {
		return false
	}
	return true
}

func Email(email string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(email) {
		return false
	}
	return true
}
