//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
)

func main() {
	opts := []entc.Option{
		entc.TemplateFiles("./tmpl/meta.tmpl"),
	}
	if err := entc.Generate("./schema", &gen.Config{
		IDType:  &field.TypeInfo{Type: field.TypeInt64},
		Target:  "./../model",
		Package: "mall-go/internal/data/model",
		Hooks: []gen.Hook{
			EnsureStructTag(),
		},
	}, opts...); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
func EnsureStructTag() gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			
			return next.Generate(g)
		})
	}
}
