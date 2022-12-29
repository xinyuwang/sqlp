package sqlp

import (
	"fmt"
	"testing"
)

func TestValueOf(t *testing.T) {

	/*
		dsn := "root:XXX@tcp(127.0.0.1:3306)/test"
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			t.Errorf(err.Error())
		}
		defer db.Close()
	*/

	type tt struct {
		AA int    `db:"A"`
		BB int    `db:"B"`
		CC string `db:"C"`
	}

	// test single

	ss := tt{
		AA: 10,
		BB: 20,
		CC: "30",
	}

	s := ValueOf(ss, "AA", "BB", "CC")
	sqlt := Query(`INSERT INTO t (%s)`, s)

	fmt.Println(sqlt.SqlStr())
	fmt.Println(sqlt.SqlArgs()...)

	/*
		_, err = db.Exec(sqlt.SqlStr(), sqlt.SqlArgs()...)
		if err != nil {
			t.Errorf(err.Error())
		}
	*/

	// test array

	ssa := []tt{
		{
			AA: 11,
			BB: 22,
			CC: "3",
		},
		{
			AA: 44,
			BB: 55,
			CC: "6",
		},
	}

	sa := ValueOf(ssa, "AA", "BB", "CC")
	sqlt = Query(`INSERT INTO t (%s)`, sa)

	fmt.Println(sqlt.SqlStr())
	fmt.Println(sqlt.SqlArgs()...)

	/*
		_, err = db.Exec(sqlt.SqlStr(), sqlt.SqlArgs()...)
		if err != nil {
			t.Errorf(err.Error())
		}
	*/

}
