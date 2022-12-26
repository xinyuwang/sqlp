package sqlp

type setPart struct {
	setStr  string
	setArgs []interface{}
}

/*
sample:
SELECT * FROM t WHERE id = (SELECT id FROM k) AND name IN (SELECT name FROM k WHERE type = 3);

let qid = sqlp.Query(`SELECT id FROM k`)
let qname = sqlp.Query(`SELECT name FROM k WHERE type = ?`. 3)

let sql = sqlp.Query(`SELECT * FROM t WHERE id = (%s) AND name IN (%s)`, qid, qname)

*/

func Set(v ...interface{}) SqlPack {

	return nil
}
