package sqlp

type sqlPack struct {
	sqlArgs []interface{}
	sqlStr  string
}

type SqlPack interface {
	SqlArgs() []any
	SqlStr() string
}

func (s *sqlPack) SqlArgs() []any {
	return s.sqlArgs
}

func (s *sqlPack) SqlStr() string {
	return s.sqlStr
}

func pack(str string, args []any) SqlPack {
	return &sqlPack{
		sqlStr:  str,
		sqlArgs: args,
	}
}
