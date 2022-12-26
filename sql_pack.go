package sqlp

type sqlPack struct {
	sqlargs []interface{}
	sqlstr  string
}

type SqlPack interface {
	Sql() (string, []interface{})
	SqlArgs() []interface{}
	SqlStr() string
}

func (s *sqlPack) SqlArgs() []interface{} {
	return s.sqlargs
}

func (s *sqlPack) SqlStr() string {
	return s.sqlstr
}

func (s *sqlPack) Sql() (string, []interface{}) {
	return s.sqlstr, s.sqlargs
}

func pack(str string, args []interface{}) SqlPack {
	return &sqlPack{
		sqlstr:  str,
		sqlargs: args,
	}
}
