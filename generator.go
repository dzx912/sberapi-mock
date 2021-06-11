package main

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type Generator struct {
	Schema *openapi3.SchemaRef
}

func (g *Generator) Generate() (interface{}, error) {
	return generate(0, g.Schema)
}

func generate(level int, schema *openapi3.SchemaRef) (interface{}, error) {
		t := schema.Value.Type

		if t == "string" {
			return "fixture-string", nil
		}

		if t == "number" {
			return 0.0, nil
		}

		if t == "integer" {
			return 0, nil
		}

		if t == "array" {
			return make([]interface{}, 0), nil
		}

		if t == "boolean" {
			return false, nil
		}

	if (t == "object" || t == "") && len(schema.Value.Properties) > 0 {
		resultMap := make(map[string]interface{})
		for k, v := range schema.Value.Properties {
			nextValue, err := generate(level + 1, v)

			if err != nil {
				return nil, err
			}

			resultMap[k] = nextValue
		}

		return resultMap, nil
	}

	return nil, fmt.Errorf("unexpected schema format")
}