package dialect

import "reflect"

// 用于兼容不同类型的数据库

type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

var dialectMap = map[string]Dialect{}

func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return
}
