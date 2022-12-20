package validators

import (
	"fmt"
	"net/mail"
	"time"
)

// returns empty string if valid or the validator
type Validate func(value any) string

/*
returns true if it is valid, false otherwise
*/
func ValidateRequired(value any) string {
	if IsNullOrEmpty(value) {
		return "field required"
	}
	return ""
}

func IsNullOrEmpty(value any) bool {
	if value == nil {
		return true
	}
	switch value.(type) {
	case string:
		if len(value.(string)) == 0 {
			return true
		}
	}
	return false
}

func ValidatorPassword(value any) string {
	if value == "" {
		return ""
	}

	stringValue, isOk := value.(string)
	if !isOk {
		return "not a string value"
	}

	if len(stringValue) < 5 {
		return "the password should have more than 5 characters"
	}
	return ""
}

func ValidateStringLength(minSize *int, maxSize *int) func(any) string {
	return func(value any) string {
		if value == nil {
			return ""
		}

		stringValue, isOk := value.(string)
		if !isOk {
			return "not a string value"
		}
		if minSize != nil {
			valueLength := len(stringValue)
			if valueLength < *minSize {
				return fmt.Sprintf("the length cannot be lower than %d", *minSize)
			}
		}
		if maxSize != nil {
			valueLength := len(stringValue)
			if valueLength > *maxSize {
				return fmt.Sprintf("the length cannot be greater than %d", *maxSize)
			}
		}

		return ""
	}
}

func DateTimeValidator(minTime *time.Time, maxTime *time.Time) func(any) string {
	return func(value any) string {
		if value == nil {
			return ""
		}
		timeValue, isOk := value.(time.Time)
		if !isOk {
			return "not a valid time format"
		}

		if timeValue.IsZero() {
			return ""
		}

		if minTime != nil {
			if timeValue.Before(*minTime) {
				return fmt.Sprintf("the date time %s cannot be before than %s", timeValue.Format(time.ANSIC), (*minTime).Format(time.ANSIC))
			}
		}
		if maxTime != nil {
			if timeValue.After(*maxTime) {
				return fmt.Sprintf("the date time %s cannot be after than %s", timeValue.Format(time.ANSIC), (*maxTime).Format(time.ANSIC))
			}
		}

		return ""
	}
}

func ValidatorEmail(value any) string {
	email, isOk := value.(string)
	if !isOk {
		return "not a string value"
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return "invalid email format"
	}

	return ""
}
