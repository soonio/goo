package schema

import (
	"go/ast"
	"goo/example/orm/dialect"
	"reflect"
)

type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FiledNames []string
	filedMap   map[string]*Field
}

func (s *Schema) GetField(name string) *Field {
	return s.filedMap[name]
}

func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modeType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modeType.Name(),
		filedMap: make(map[string]*Field),
	}
	for i := 0; i < modeType.NumField(); i++ {
		p := modeType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FiledNames = append(schema.FiledNames, p.Name)
			schema.filedMap[p.Name] = field
		}
	}
	return schema
}
