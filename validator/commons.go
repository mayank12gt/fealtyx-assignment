package validator

import "regexp"

var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func ValidateIntegerRange(key, low, high int) bool {
	return key >= low && key <= high
}

func ValidateEmail(key string) bool {
	email_regex := regexp.MustCompile(emailRegex)

	return email_regex.MatchString(key)
}
