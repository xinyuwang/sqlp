package sqlp

import (
	"strings"
)

/*
Expend slice into a set of elements
*/
func In(v ...any) SqlPack {

	arrStr := []string{}
	arrArgs := []interface{}{}

	for _, val := range v {

		arrStr = append(arrStr, "?")
		arrArgs = append(arrArgs, val)
	}

	str := strings.Join(arrStr, ", ")

	return pack(str, arrArgs)
}
