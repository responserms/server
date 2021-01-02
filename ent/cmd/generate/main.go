package main

import (
	"log"

	"github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
	"github.com/facebookincubator/ent-contrib/entgql"
	"github.com/responserms/server/ent/schema/templates"
)

func main() {
	entTemplates := []*gen.Template{}

	// add the entgql templates
	entTemplates = append(entTemplates, entgql.AllTemplates...)

	// add our internal templates
	entTemplates = append(entTemplates, templates.InternalTemplates...)

	err := entc.Generate("./ent/schema", &gen.Config{
		Templates: entTemplates,
	})

	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
