package main

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"text/template"

	"github.com/tabbed/sqlc-go/codegen"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const crudTmpl = `-- Code generated by sqlc-crud. DO NOT EDIT.
{{ $table := .Rel.Name }}
-- name: Insert{{$table | pascal}} :execresult 
INSERT INTO {{$table}} (
	{{ range $i, $column := .Columns -}}
		{{ if ne $column.Name "id" }}{{if ne $i 1}},{{end}} {{$column.Name}}{{ end }}
	{{- end }}
) VALUES (
	{{ range $i, $column := .Columns -}}
		{{ if ne $column.Name "id"}}{{if ne $i 1}},{{end}} ?{{ end }}
	{{- end }}
);
{{ $table := .Rel.Name }}
-- name: Update{{$table | pascal}} :exec
UPDATE {{$table}} SET
	{{ range $i, $column := .Columns -}}
		{{ if ne $column.Name "id" }}{{if ne $i 1}},{{end}} {{$column.Name}} = ?{{ end }}
	{{- end }}
WHERE id = ?;

-- name: Get{{$table | pascal}} :one
SELECT * FROM {{$table}} WHERE id = ?;
{{- range $i, $column := .Columns -}}
{{- if eq $column.Name "guid" }}

-- name: Get{{$table | pascal}}ByGUID :one
SELECT * FROM {{$table}} WHERE guid = ?;
{{ end -}}
{{- end }}
-- name: Delete{{$table | pascal}} :exec
DELETE FROM {{$table}} WHERE id = ?;
`

func toSnakeCase(s string) string {
	return strings.ToLower(s)
}

func toCamelCase(s string) string {
	return toCamelInitCase(s, false)
}

func toPascalCase(s string) string {
	return toCamelInitCase(s, true)
}

func toCamelInitCase(name string, initUpper bool) string {
	out := ""
	caser := cases.Title(language.English)
	for i, p := range strings.Split(name, "_") {
		if !initUpper && i == 0 {
			out += p
			continue
		}
		if p == "id" {
			out += "ID"
		} else if p == "guid" {
			out += "GUID"
		} else {
			out += caser.String(p)
		}
	}
	return out
}

func main() {
	codegen.Run(generate)
}

func generate(_ context.Context, req *codegen.Request) (*codegen.Response, error) {
	funcMap := map[string]interface{}{
		"snake":  toSnakeCase,
		"pascal": toPascalCase,
		"camel":  toCamelCase,
	}

	tmpl := template.Must(template.New("crud").Funcs(funcMap).Parse(crudTmpl))

	res := codegen.CodeGenResponse{
		Files: []*codegen.File{},
	}

	for _, schema := range req.Catalog.Schemas {
		if schema.Name == req.Catalog.DefaultSchema {
			for _, table := range schema.Tables {
				if len(table.Columns) > 0 && table.Columns[0].Name == "id" {
					var b bytes.Buffer
					err := tmpl.Execute(&b, *table)
					if err != nil {
						return nil, err
					}
					file := codegen.File{
						Name:     fmt.Sprintf("%s.crud.sql", table.Rel.Name),
						Contents: b.Bytes(),
					}
					res.Files = append(res.Files, &file)
				}
			}
		}
	}

	return &res, nil
}
