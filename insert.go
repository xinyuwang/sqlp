package sqlp

import (
	"fmt"
	"reflect"
	"strings"
)

/*
Transform a struct or slice of struct into INSERT expression.
v is a struct or slice
fields include the name of struct fields.
*/
func ValueOf(v any, fields ...string) SqlPack {

	vof := reflect.ValueOf(v)
	for vof.Kind() == reflect.Pointer {
		vof = vof.Elem()
	}

	// to array
	var stuArr []reflect.Value
	if vof.Kind() == reflect.Slice {
		// vof.Len() should above 1
		if vof.Len() == 0 {
			panic("param 1 should not be empty slice")
		}

		for i := 0; i < vof.Len(); i++ {
			item := vof.Index(i)
			for item.Kind() == reflect.Pointer {
				item = item.Elem()
			}

			stuArr = append(stuArr, item)
		}

	} else if vof.Kind() == reflect.Struct {
		stuArr = append(stuArr, vof)
	}

	// fetch db tags into columns
	var colArr []string
	tof := stuArr[0].Type()
	for i := 0; i < len(fields); i++ {

		stuMem := fields[i]

		// process type
		stuField, ok := tof.FieldByName(stuMem)
		if !ok {
			panic(fmt.Sprintf("struct filed %s not found", stuMem))
		}

		tagVal := stuField.Tag.Get("db")
		colArr = append(colArr, tagVal)
	}

	// fetch all insert values

	var arrStr []string
	var arrArgs []interface{}
	for _, stu := range stuArr {

		var valArr []string

		for i := 0; i < len(fields); i++ {

			valArr = append(valArr, "?")

			stuMem := fields[i]
			fieldVal := stu.FieldByName(stuMem)
			arrArgs = append(arrArgs, fieldVal.Interface())
		}

		arrStr = append(arrStr, strings.Join(valArr, ", "))
	}

	str := fmt.Sprintf("%s) VALUES (%s", strings.Join(colArr, ", "), strings.Join(arrStr, "), ("))

	return pack(str, arrArgs)
}
