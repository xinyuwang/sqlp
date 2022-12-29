package sqlp

import (
	"fmt"
	"testing"
)

func TestIn(t *testing.T) {

	//SELECT * FROM t WHERE id IN (1, 2, 3)

	s := In(1, 2, 3)
	sqlt := Query(`SELECT * FROM t WHERE id IN (%s)`, s)

	fmt.Println(sqlt.SqlStr())

	fmt.Println(sqlt.SqlArgs()...)

}
