package util

import (
	"regexp"
)

func RegexpParentheses(str string) (string, error) {
	reg, err := regexp.Compile(".*[\\(（]{1}(.*)[\\)）]{1}.*")
	if err != nil {
		return "", err
	}
	sublist := reg.FindStringSubmatch(str)
	if len(sublist) < 2 {
		return "", err
	}
	return sublist[1], nil
}
