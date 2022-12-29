package sqlp

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {

	tt := 3

	qid := Query(`SELECT id FROM k`)
	qname := Query(`SELECT name FROM k WHERE type = ?`, tt)
	sqlt := Query(`SELECT * FROM t WHERE id = (%s) AND name IN (%s)`, qid, qname)

	fmt.Println(sqlt.SqlStr())
	fmt.Println(sqlt.SqlArgs()...)
}
