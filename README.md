# sqlp
SQL Pack
简单的SQL语句拼接库

## 主要功能

1. 子查询复用

通过复用查询，可以将复杂的SQL语句化简，优化代码表达。具体用法如下：

```golang
// SELECT * FROM t WHERE id = (SELECT id FROM k) AND name IN (SELECT name FROM k WHERE type = 3)

tt := 3
qname := sqlp.Query(`SELECT name FROM k WHERE type = ?`, tt)
qid := sqlp.Query(`SELECT id FROM k`)

// 复用查询
sqlt := sqlp.Query(`SELECT * FROM t WHERE id = (%s) AND name IN (%s)`, qid, qname)

// 由 Go-Mysql 运行
_, err = db.Exec(sqlt.SqlStr(), sqlt.SqlArgs()...)
```

2. 查询条件IN切片展开

```golang
// SELECT * FROM t WHERE id IN (1, 2, 3)

s := sqlp.In(1, 2, 3)

sqlt := sqlp.Query(`SELECT * FROM t WHERE id IN (%s)`, s)

// 由 Go-Mysql 运行
_, err = db.Exec(sqlt.SqlStr(), sqlt.SqlArgs()...)
```

3. Update更新数据同侧配置

直接写 `UPDATE SET Column = Value` 语句的话，前面Column 与 值参数不在同侧，需要按顺序写好，容易错位。通过使用Set函数，可以将Column 和 Value 交替书写。具体用法如下：

```golang
// UPDATE t SET A=1, B=2, C='3' WHERE id=4

a := 1
b := 1
c := '3'
id := 4

s := sqlp.Set(`A`, a, `B`, b, `C`, c)
sqlt := sqlp.Query(`UPDATE t SET %s WHERE id = ?`, s, id)

// 由 Go-Mysql 运行
_, err = db.Exec(sqlt.SqlStr(), sqlt.SqlArgs()...)
```


4. Update支持结构体反射更新

支持使用结构体反射，通过`SetOf`函数，指定需要更新的结构体成员名称。具体用法如下：

```golang
// UPDATE t SET A = 1, B = 2 WHERE id = 4
type tt struct {
	AA int    `db:"A"`
	BB int    `db:"B"`
	CC string `db:"C"`
}

ss := &tt{
	AA: 1,
	BB: 2,
	CC: "3",
}

id := 4

// 后面参数为结构体成员的名称
s := sqlp.SetOf(ss, "AA", "BB")
sqlt := sqlp.Query(`UPDATE t SET %s WHERE id = ?`, s, id)

// 由 Go-Mysql 运行
_, err = db.Exec(sqlt.SqlStr(), sqlt.SqlArgs()...)
```

5. Insert支持结构体反射插入

支持使用结构体反射，通过`ValueOf`函数，指定需要更新的结构体成员名称，并根据输入生成单个或批量插入的语句。具体用法如下：

```golang
// 单条插入
// INSERT INTO t (A, B, C) VALUES (10, 20, 30)

type tt struct {
	AA int    `db:"A"`
	BB int    `db:"B"`
	CC string `db:"C"`
}

ss := tt{
	AA: 10,
	BB: 20,
	CC: "30",
}

// 单个结构体，后面参数为结构体成员的名称
s := sqlp.ValueOf(ss, "AA", "BB", "CC")
sqlt := sqlp.Query(`INSERT INTO t (%s)`, s)

// 由 Go-Mysql 运行
_, err = db.Exec(sqlt.SqlStr(), sqlt.SqlArgs()...)

// 批量插入
// INSERT INTO t (A, B, C) VALUES (10, 20, 30), (40, 50, 60) 
    
ssa := []tt{
	{
		AA: 10,
		BB: 20,
		CC: "30",
	},
	{
		AA: 40,
		BB: 50,
		CC: "60",
	},
}

// 结构体切片，后面参数为结构体成员的名称
sa := sqlp.ValueOf(ssa, "AA", "BB", "CC")
sqlt = sqlp.Query(`INSERT INTO t (%s)`, sa)

// 由 Go-Mysql 运行
_, err = db.Exec(sqlt.SqlStr(), sqlt.SqlArgs()...)
```