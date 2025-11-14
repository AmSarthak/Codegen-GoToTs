package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

var typeMap = map[string]string{
	"string":  "string",
	"int":     "number",
	"int64":   "number",
	"float64": "number",
	"bool":    "boolean",
}

func toTSType(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return typeMap[t.Name]
	case *ast.ArrayType:
		return toTSType(t.Elt) + "[]"
	default:
		return "any"
	}
}

func main() {
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "backend/models/DataModel.go", nil, parser.ParseComments)

	var output strings.Builder

	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}
			output.WriteString("export interface " + typeSpec.Name.Name + " {\n")

			for _, field := range structType.Fields.List {
				tsType := toTSType(field.Type)
				for _, name := range field.Names {
					output.WriteString(fmt.Sprintf("  %s: %s;\n", name.Name, tsType))
				}
			}

			output.WriteString("}\n\n")
		}
	}

	os.WriteFile("models.ts", []byte(output.String()), 0644)
}
