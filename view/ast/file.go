package ast

import (
	"bytes"
	"fmt"

	"github.com/things-go/ormat/pkg/consts"
)

// File a file
type File struct {
	Filename        string            // file name
	PackageName     string            // package name
	Imports         map[string]string // import package
	Structs         []*Struct         // struct list in file
	IsOutColumnName bool              // is out column name
	ProtobufPackage string
	ProtobufOptions map[string]string
}

// AddImport Add import by type
func (p *File) AddImport(imp string) *File {
	if p.Imports == nil {
		p.Imports = make(map[string]string)
	}
	p.Imports[imp] = imp
	return p
}

// SetOutColumnName set out column name
func (p *File) SetOutColumnName(b bool) *File {
	p.IsOutColumnName = b
	return p
}

// AddStruct Add a structure
func (p *File) AddStruct(st *Struct) *File {
	p.Structs = append(p.Structs, st)
	return p
}

// Build the file data
func (p *File) Build() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("// Code generated by ormat. DO NOT EDIT.\n")
	buf.WriteString("// version: " + consts.Version + "\n")
	buf.WriteString("\n")
	buf.WriteString("package" + delimTab + p.PackageName + delimLF)

	// auto add import
	for _, v := range p.Structs {
		for _, v1 := range v.StructFields {
			if v2, ok := ImportsHeads[v1.FieldType]; ok && v2 != "" {
				p.AddImport(v2)
			}
		}
	}

	// add imports
	if len(p.Imports) > 0 {
		buf.WriteString("import (" + delimLF)
		for _, v := range p.Imports {
			buf.WriteString(v + delimLF)
		}
		buf.WriteString(")" + delimLF)
	}

	// add struct
	for _, v := range p.Structs {
		buf.WriteString(v.Build() + delimLF)
		// add table name function
		buf.WriteString(v.BuildTableNameTemplate() + delimLF)
		if p.IsOutColumnName {
			buf.WriteString(delimLF)
			buf.WriteString(v.BuildColumnNameTemplate())
		}
		buf.WriteString(v.BuildProtobufEnumMappingTemplate())

		buf.WriteString(v.BuildProtobufTemplate() + delimLF)
	}
	return buf.Bytes()
}

// BuildSQL sql in one file
func (p *File) BuildSQL() []byte {
	buf := bytes.Buffer{}

	// add struct
	for _, v := range p.Structs {
		buf.WriteString(v.BuildSQL() + delimLF)
	}
	return buf.Bytes()
}

func (p *File) BuildProtobufEnum() []byte {
	body := p.BuildProtobufEnumBody()
	if len(body) == 0 {
		return []byte{}
	}
	header := BuildRawProtobufEnumHeader(p.ProtobufPackage, p.ProtobufOptions)
	return append(header, body...)
}

func (p *File) BuildProtobufEnumBody() []byte {
	buf := bytes.Buffer{}

	for _, v := range p.Structs {
		content := v.BuildProtobufEnumTemplate()
		if len(content) > 0 {
			buf.WriteString(content + delimLF)
		}
	}
	return buf.Bytes()
}

func BuildRawProtobufEnumHeader(protobufPackage string, protobufOptions map[string]string) []byte {
	buf := bytes.Buffer{}
	buf.WriteString("// Code generated by ormat. DO NOT EDIT.\n")
	buf.WriteString("// version: " + consts.Version + "\n")
	buf.WriteString("\n")
	buf.WriteString(`syntax = "proto3";` + delimLF + delimLF)
	buf.WriteString(fmt.Sprintf("package %s;", protobufPackage) + delimLF)
	for k, v := range protobufOptions {
		buf.WriteString(fmt.Sprintf(`option %s = "%s";`, k, v) + delimLF)
	}
	return buf.Bytes()
}
