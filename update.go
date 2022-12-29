package sqlp

import (
	"fmt"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

/*
Build a Set struct for update table.
v includes pairs of column name and value.
*/
func Set(v ...any) SqlPack {

	arrStr := []string{}
	arrArgs := []interface{}{}

	for i := 0; i < len(v); i += 2 {

		s := fmt.Sprintf("%s = ?", v[i])
		arrStr = append(arrStr, s)

		arrArgs = append(arrArgs, v[i+1])
	}

	str := strings.Join(arrStr, ", ")

	return pack(str, arrArgs)
}

/*
Transform a struct or slice of struct into UPDATE expression.
v is a struct
fields include the name of struct fields
*/
func SetOf(v any, fields ...string) SqlPack {

	arrStr := []string{}
	arrArgs := []interface{}{}

	vof := reflect.ValueOf(v)
	for vof.Kind() == reflect.Pointer {
		vof = vof.Elem()
	}
	tof := vof.Type()

	for i := 0; i < len(fields); i++ {

		stuMem := fields[i]

		// process type
		stuField, ok := tof.FieldByName(stuMem)
		if !ok {
			panic(fmt.Sprintf("struct filed %s not found", stuMem))
		}

		tagVal := stuField.Tag.Get("db")
		s := fmt.Sprintf("%s = ?", tagVal)
		arrStr = append(arrStr, s)

		// process value
		v := vof.FieldByName(stuMem)

		arrArgs = append(arrArgs, v.Interface())
	}

	str := strings.Join(arrStr, ", ")

	return pack(str, arrArgs)
}
