package sqlp

type SqlType int

const (
	TSelect SqlType = 101
	TInsert SqlType = 102
	TUpdate SqlType = 103
	TDelete SqlType = 104
)
