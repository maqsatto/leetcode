package leet20

import (
	"strings"
)

// not solved
func isValid(s string) bool {

	valid := false
	str := strings.Split(s, "")
	for _, v := range str {
		v = strings.Replace(v, "*", "", -1)
	}
	return valid
}
