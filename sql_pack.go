package sqlp

type sqlPack struct {
	sqlargs []interface{}
	sqlstr  string
}

type Package interface {
	SqlArgs() []interface{}
	SqlStr() string
}

func (s *sqlPack) SqlArgs() []interface{} {
	return s.sqlargs
}

func (s *sqlPack) SqlStr() string {
	return s.sqlstr
}

func Pack(str string, args []interface{}) Package {
	return &sqlPack{
		sqlstr:  str,
		sqlargs: args,
	}
}
