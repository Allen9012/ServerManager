package file

import (
	"regexp"
)

func isRegixValid(regStr string) error {
	_, err := regexp.Compile(regStr)
	return err
}
