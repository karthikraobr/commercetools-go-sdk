package main

import (
	"log"

	"github.com/dave/jennifer/jen"
)

// Entry point to generate code object for a specific RamlType
func generateTypes(objects []RamlType) {
	items := map[string][]*RamlType{}

	for i := range objects {
		item := objects[i]
		items[item.Package] = append(items[item.Package], &item)
	}

	for pkg, packageObjects := range items {
		filename := generateFilename(pkg, "")
		log.Printf("Writing commercetools/%s\n", filename)
		f := jen.NewFile("commercetools")
		f.HeaderComment("Automatically generated, do not edit")

		enumObjects := map[string]jen.Code{}
		mapObjects := map[string]jen.Code{}
		stringObjects := map[string]jen.Code{}
		structObjects := map[string]jen.Code{}

		for _, object := range packageObjects {
			var stmt *jen.Statement
			key := object.CodeName

			// Just to keep sorting order (bit of a hack)
			if object.isInterface() {
				key = "Abstract" + key
			}

			if object.asMap {
				stmt = generateMap(*object)
				mapObjects[key] = stmt
			} else if len(object.EnumValues) > 0 {
				stmt := generateEnum(*object)
				enumObjects[key] = stmt
			} else if object.TypeName == "string" {
				stmt := generateString(*object)
				stringObjects[key] = stmt
			} else {
				stmt = generateStruct(*object)
				structObjects[key] = stmt
			}
		}

		addCodeObjects(f, enumObjects)
		addCodeObjects(f, stringObjects)
		addCodeObjects(f, mapObjects)
		addCodeObjects(f, structObjects)

		err := f.Save("commercetools/types_" + filename)
		if err != nil {
			panic(err)
		}
	}
}
