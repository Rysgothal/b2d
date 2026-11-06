package validation

import "errors"

func Binary(value string) error {
	if value == "" {
		return errors.New("binary value cannot be empty")
	}

	if len(value) > 8 {
		return errors.New("binary value cannot be longer than 8 digits")
	}

	for _, digit := range value {
		if digit != '0' && digit != '1' {
			return errors.New("value must contain only 0 and 1")
		}
	}

	return nil
}
