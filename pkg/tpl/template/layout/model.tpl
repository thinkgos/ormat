// Code generated by ormat. DO NOT EDIT.
// version: {{.Version}}

package {{.PackageName}}

{{if .Imports}}
import (
{{- range $k, $v := .Imports}}
    {{$k}}
{{- end}}
)
{{end}}

{{- $hasColumn := .HasColumn}}
{{- $hasHelper := .HasHelper}}
{{- $hasAssist := .HasAssist}}
{{- range $e := .Structs}}
// {{$e.StructName}} {{$e.StructComment}}
type {{$e.StructName}} struct {
{{- range $field := $e.StructFields}}
    {{$field.FieldName}} {{$field.FieldType}} {{if $field.FieldTag}}`{{$field.FieldTag}}`{{end}} {{if $field.FieldComment}}// {{$field.FieldComment}}{{end}}
{{- end}}
}

// TableName implement schema.Tabler interface
func (*{{$e.StructName}}) TableName() string {
	return "{{$e.TableName}}"
}

{{- $tableName := $e.TableName}}
{{- if $hasColumn}}
// Select{{$e.StructName}} database column name.
var Select{{$e.StructName}} = []string {
{{- range $field := $e.StructFields}}
	{{- if $field.IsTimestamp}}
	{{- if $field.IsNullable}}
	{{if $field.IsSkipColumn}}// {{end}}"IFNULL(UNIX_TIMESTAMP(`{{$tableName}}`.`{{$field.ColumnName}}`), 0) AS `{{$field.ColumnName}}`",
	{{- else}}
	{{if $field.IsSkipColumn}}// {{end}}"UNIX_TIMESTAMP(`{{$tableName}}`.`{{$field.ColumnName}}`) AS `{{$field.ColumnName}}`",
	{{- end}}
	{{- else}}
	{{if $field.IsSkipColumn}}// {{end}}"`{{$tableName}}`.`{{$field.ColumnName}}`",
	{{- end}}
{{- end}}
}
{{- end}}

{{- if $hasAssist}}

const (
	// hold model `{{$e.StructName}}` table name
	xx_{{$e.StructName}}_TableName = "{{$e.TableName}}"
	// hold model `{{$e.StructName}}` column name
{{- range $field := $e.StructFields}}
    xx_{{$e.StructName}}_{{$field.FieldName}} = "{{$field.ColumnName}}"
{{- end}}
)

var xxx_{{$e.StructName}}_Model = new_X_{{$e.StructName}}(xx_{{$e.StructName}}_TableName)
var xxx_{{$e.StructName}}_ActiveModel = new_X_{{$e.StructName}}("")

type {{$e.StructName}}Impl_x struct {
	// private fields
	xTableName string 

	ALL assist.Asterisk
{{- range $field := $e.StructFields}}
    {{$field.FieldName}} assist.{{$field.AssistType}}
{{- end}}
}

// X_{{$e.StructName}} model with TableName `{{$e.TableName}}`.
func X_{{$e.StructName}}() {{$e.StructName}}Impl_x {
	return xxx_{{$e.StructName}}_Model
}

// X_Active_{{$e.StructName}} active model without TableName.
func X_Active_{{$e.StructName}}() {{$e.StructName}}Impl_x {
	return xxx_{{$e.StructName}}_ActiveModel
}

func new_X_{{$e.StructName}}(tableName string) {{$e.StructName}}Impl_x {
	return {{$e.StructName}}Impl_x{
		xTableName: tableName,

		ALL:  assist.NewAsterisk(tableName),

	{{range $field := $e.StructFields}}
		{{$field.FieldName}}: assist.New{{$field.AssistType}}(tableName, xx_{{$e.StructName}}_{{$field.FieldName}}),
	{{- end}}			
	}
}

func New_X_{{$e.StructName}}(xTableName string) {{$e.StructName}}Impl_x {
	if xTableName == "" {
		return xxx_{{$e.StructName}}_ActiveModel
	} else {
		return new_X_{{$e.StructName}}(xTableName)
	}
}

// TableName hold table name when call New_X_{{$e.StructName}} or {{$e.StructName}}Impl_x.As.
func (x *{{$e.StructName}}Impl_x) X_TableName() string {
	return x.xTableName
}

func (*{{$e.StructName}}Impl_x) As(alias string) {{$e.StructName}}Impl_x {
	return New_X_{{$e.StructName}}(alias)
}

func (*{{$e.StructName}}Impl_x) X_Model() *{{$e.StructName}} {
	return &{{$e.StructName}}{}
}

func (*{{$e.StructName}}Impl_x) Xc_Model() assist.Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Model(&{{$e.StructName}}{})
	}
}

// TableName hold model `{{$e.StructName}}` table name returns `{{$e.TableName}}`.
func (x *{{$e.StructName}}Impl_x) TableName() string {
	return xx_{{$e.StructName}}_TableName
}
{{- range $field := $e.StructFields}}
// Field_{{$field.FieldName}} hold model `{{$e.StructName}}` column name.
// if prefixes not exist returns `{{$field.ColumnName}}`, others `{prefixes[0]}_{{$field.ColumnName}}`
func (*{{$e.StructName}}Impl_x) Field_{{$field.FieldName}}(prefixes ...string) string {
	if len(prefixes) > 0 {
		return prefixes[0] + "_" + xx_{{$e.StructName}}_{{$field.FieldName}}
	}
	return xx_{{$e.StructName}}_{{$field.FieldName}}
}
{{- end}}

func X_Select{{$e.StructName}}(prefixes ...string) []assist.Expr {
	x := &xxx_{{$e.StructName}}_Model
	if len(prefixes) > 0 {
		prefix := prefixes[0] + "_"
		return []assist.Expr{
	{{- range $field := $e.StructFields}}
		{{if $field.IsSkipColumn}}// {{end}}x.{{$field.FieldName}}{{- if $field.IsTimestamp}}.UnixTimestamp(){{- if $field.IsNullable}}.IfNull(0){{- end}}{{- end}}.As(prefix + xx_{{$e.StructName}}_{{$field.FieldName}}),
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

func Xc_Select{{$e.StructName}}(prefixes ...string) assist.Condition {
	return assist.Select(X_Select{{$e.StructName}}(prefixes...)...)
}

{{- end}}

{{- if $hasHelper}}
/* protobuf field helper
// {{$e.StructName}} {{.StructComment}}
message {{$e.StructName}} {
{{- range $index, $field := $e.ProtoMessageFields}}
  {{- if $field.FieldComment}}
  // {{$field.FieldComment}}
  {{- end}}
  {{$field.FieldDataType}} {{$field.FieldName}} = {{add $index 1}} {{- if $field.FieldAnnotation}} {{$field.FieldAnnotation}} {{- end}};
{{- end}}
}
*/
{{- end}}
{{- end}}

