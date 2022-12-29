package sqlp

import "fmt"

/*
Build a query with string and corresponding args.
*/
func Query(str string, v ...any) SqlPack {

	arrStr := []interface{}{}
	arrArgs := []interface{}{}

	for _, val := range v {

		if p, ok := val.(SqlPack); ok {
			arrStr = append(arrStr, p.SqlStr())
			arrArgs = append(arrArgs, p.SqlArgs()...)
		} else {
			arrArgs = append(arrArgs, val)
		}
	}

	// check the %s match if needed;
	// for now, the user checks it.
	str = fmt.Sprintf(str, arrStr...)

	return pack(str, arrArgs)
}
