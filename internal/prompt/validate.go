package prompt

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func ValidateMinLength(min int) promptui.ValidateFunc {
	return func(value string) error {
		if len(value) < min {
			return fmt.Errorf("Length must be greater than %d", min)
		}
		return nil
	}
}

func ValidateYesNo() promptui.ValidateFunc {
	return func(value string) error {
		switch value[0] {
		case 'Y', 'y', 'N', 'n':
			return nil
		default:
			return fmt.Errorf("Must be y or n")
		}
	}
}
