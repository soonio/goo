package clause

import "strings"

type Type int

const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
)

type Clause struct {
	sql    map[Type]string
	sqlVar map[Type][]interface{}
}

func (c *Clause) Set(name Type, values ...interface{}) {
	if c.sql == nil {
		c.sql = make(map[Type]string)
		c.sqlVar = make(map[Type][]interface{})
	}
	sql, values := generators[name](values...)
	c.sql[name] = sql
	c.sqlVar[name] = values
}

// 构建sql及其参数参数
func (c *Clause) Build(order ...Type) (string, []interface{}) {
	var sqls []string
	var vars []interface{}
	for _, order = range order {
		if sql, ok := c.sql[order]; ok {
			sqls = append(sqls, sql)
			vars = append(vars, c.sqlVar[order]...)
		}
	}
	return strings.Join(sqls, " "), vars
}
