package sqlp

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {

	// UPDATE t SET A=1, B=2, C='3' WHERE id=4

	a := 1
	b := 1
	c := '3'
	id := 4

	s := Set(`A`, a, `B`, b, `C`, c)
	sqlt := Query(`UPDATE t SET %s WHERE id = ?`, s, id)

	fmt.Println(sqlt.SqlStr())

	fmt.Println(sqlt.SqlArgs()...)

}

func TestSetOf(t *testing.T) {

	/*
		dsn := "root:XXX@tcp(127.0.0.1:3306)/test"
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			t.Errorf(err.Error())
		}
		defer db.Close()
	*/

	// UPDATE t SET A=1, B=2, C='3' WHERE id=4

	type tt struct {
		AA int    `db:"A"`
		BB int    `db:"B"`
		CC string `db:"C"`
	}

	ss := &tt{
		AA: 44,
		BB: 88,
		CC: "3333",
	}

	id := 4

	s := SetOf(ss, "AA", "BB")
	sqlt := Query(`UPDATE t SET %s WHERE id = ?`, s, id)

	fmt.Println(sqlt.SqlStr())

	fmt.Println(sqlt.SqlArgs()...)

	/*
		_, err = db.Exec(sqlt.SqlStr(), sqlt.SqlArgs()...)
		if err != nil {
			t.Errorf(err.Error())
		}
	*/

}
