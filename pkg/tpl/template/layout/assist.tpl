// Code generated by ormat. DO NOT EDIT.
// version: {{.Version}}

package {{.PackageName}}

import (
    assist "github.com/things-go/gorm-assist"
	"gorm.io/gorm"
)

{{- range $e := .Structs}}

const (
	// hold model `{{$e.StructName}}` table name
	xx_{{$e.StructName}}_TableName = "{{$e.TableName}}"
	// hold model `{{$e.StructName}}` column name
{{- range $field := $e.StructFields}}
    xx_{{$e.StructName}}_{{$field.FieldName}} = "{{$field.ColumnName}}"
{{- end}}
	// hold model `{{$e.StructName}}` column name with table name(`{{$e.TableName}}`) prefix
{{- range $field := $e.StructFields}}
    xx_{{$e.StructName}}_{{$field.FieldName}}_WithTableName = xx_{{$e.StructName}}_TableName + "_" + xx_{{$e.StructName}}_{{$field.FieldName}}
{{- end}}
)

var xxx_{{$e.StructName}}_Native_Model = new_X_{{$e.StructName}}("")
var xxx_{{$e.StructName}}_Model = new_X_{{$e.StructName}}(xx_{{$e.StructName}}_TableName)

type {{$e.StructName}}_Active struct {
	// private fields
	xTableName string 

	ALL assist.Asterisk
{{- range $field := $e.StructFields}}
    {{$field.FieldName}} assist.{{$field.AssistType}}
{{- end}}
}

// X_Native_{{$e.StructName}} native model without TableName.
func X_Native_{{$e.StructName}}() {{$e.StructName}}_Active {
	return xxx_{{$e.StructName}}_Native_Model
}

// X_{{$e.StructName}} model with TableName `{{$e.TableName}}`.
func X_{{$e.StructName}}() {{$e.StructName}}_Active {
	return xxx_{{$e.StructName}}_Model
}

func new_X_{{$e.StructName}}(xTableName string) {{$e.StructName}}_Active {
	return {{$e.StructName}}_Active{
		xTableName: xTableName,

		ALL:  assist.NewAsterisk(xTableName),
	{{range $field := $e.StructFields}}
		{{$field.FieldName}}: assist.New{{$field.AssistType}}(xTableName, xx_{{$e.StructName}}_{{$field.FieldName}}),
	{{- end}}			
	}
}

// New_X_{{$e.StructName}} new instance.
func New_X_{{$e.StructName}}(xTableName string) {{$e.StructName}}_Active {
	switch xTableName {
	case "":
		return xxx_{{$e.StructName}}_Native_Model
	case xx_{{$e.StructName}}_TableName:
		return xxx_{{$e.StructName}}_Model
	default:
		return new_X_{{$e.StructName}}(xTableName)
	}
}

// As alias
func (*{{$e.StructName}}_Active) As(alias string) {{$e.StructName}}_Active {
	return New_X_{{$e.StructName}}(alias)
}

// X_TableName hold table name when call New_X_{{$e.StructName}} or {{$e.StructName}}_Active.As that you defined.
func (x *{{$e.StructName}}_Active) X_TableName() string {
	return x.xTableName
}

// X_Model model
func (*{{$e.StructName}}_Active) X_Model() *{{$e.StructName}} {
	return &{{$e.StructName}}{}
}

// X_Executor new entity executor which suggest use only once.
func (*{{$e.StructName}}_Active) X_Executor(db *gorm.DB) *assist.Executor[{{$e.StructName}}] {
	return assist.NewExecutor[{{$e.StructName}}](db)
}

// TableName hold model `{{$e.StructName}}` table name returns `{{$e.TableName}}`.
func (*{{$e.StructName}}_Active) TableName() string {
	return xx_{{$e.StructName}}_TableName
}
{{- range $field := $e.StructFields}}
// Field_{{$field.FieldName}} hold model `{{$e.StructName}}` column name.
// if prefixes not exist returns `{{$field.ColumnName}}`, others `{prefixes[0]}_{{$field.ColumnName}}`
func (*{{$e.StructName}}_Active) Field_{{$field.FieldName}}(prefixes ...string) string {
	if len(prefixes) == 0 {
		return xx_{{$e.StructName}}_{{$field.FieldName}}
	}
	if prefixes[0] == xx_{{$e.StructName}}_TableName {
		return xx_{{$e.StructName}}_{{$field.FieldName}}_WithTableName
	}
	return prefixes[0] + "_" + xx_{{$e.StructName}}_{{$field.FieldName}}
}
{{- end}}

func SelectNative{{$e.StructName}}() []assist.Expr {
	x := &xxx_{{$e.StructName}}_Native_Model
	return []assist.Expr{
{{- range $field := $e.StructFields}}
		x.{{$field.FieldName}},
{{- end}}
	}
}

func x_Select{{$e.StructName}}(x *{{$e.StructName}}_Active, prefixes ...string) []assist.Expr {
	if len(prefixes) > 0 {
		return []assist.Expr{
		{{- range $field := $e.StructFields}}
			{{if $field.IsSkipColumn}}// {{end}}x.{{$field.FieldName}}{{- if $field.IsTimestamp}}.UnixTimestamp(){{- if $field.IsNullable}}.IfNull(0){{- end}}{{- end}}.As(x.Field_{{$field.FieldName}}(prefixes...)),
		{{- end}}
		}
	} else {
		return []assist.Expr{
		{{- range $field := $e.StructFields}}
			{{- if $field.IsTimestamp}}
			{{if $field.IsSkipColumn}}// {{end}}x.{{$field.FieldName}}.UnixTimestamp(){{- if $field.IsNullable}}.IfNull(0){{- end}}.As(xx_{{$e.StructName}}_{{$field.FieldName}}),
			{{- else}}
			{{if $field.IsSkipColumn}}// {{end}}x.{{$field.FieldName}},
			{{- end}}
		{{- end}}
		}
	}
}

// X_Native_Select{{$e.StructName}} select field use use X_Native_{{$e.StructName}}().
func X_Native_Select{{$e.StructName}}() []assist.Expr {
	return x_Select{{$e.StructName}}(&xxx_{{$e.StructName}}_Native_Model)
}

// X_Select{{$e.StructName}} select fields use X_{{$e.StructName}}().
func X_Select{{$e.StructName}}(prefixes ...string) []assist.Expr {
	return x_Select{{$e.StructName}}(&xxx_{{$e.StructName}}_Model, prefixes...)
}

{{- end}}