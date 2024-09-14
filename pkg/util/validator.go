package util

import "fmt"

func ValidateTitle(value string) error {
	if value == "" {
		return fmt.Errorf("title cannot be empty")
	}

	return nil
}
