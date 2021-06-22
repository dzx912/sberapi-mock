package oapi

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type Generator struct {
	Schema *openapi3.SchemaRef
}

func NewGenerator(schema *openapi3.SchemaRef) *Generator {
	return &Generator{ Schema: schema }
}

func (g *Generator) Generate() (interface{}, error) {
	return generate(0, g.Schema)
}

func generate(level int, schema *openapi3.SchemaRef) (interface{}, error) {
		t := schema.Value.Type

		if t == "string" {
			if schema.Value.Example != nil {
				return schema.Value.Example, nil
			}
			return "fixture-string", nil
		}

		if t == "number" {
			if schema.Value.Example != nil {
				return float64(schema.Value.Example.(float64)), nil
			}
			return float64(0.0), nil
		}

		if t == "integer" {
			if schema.Value.Example != nil {
				n := int64(schema.Value.Example.(float64)) 
				
				return n, nil
			}

			return int64(0), nil
		}

		if t == "array" {
			item, err := generate(level + 1, schema.Value.Items)

			if err != nil {
				return make([]interface{}, 0), nil
			}

			return []interface{}{ item }, nil
		}

		if t == "boolean" {
			if schema.Value.Example != nil {
				return schema.Value.Example, nil
			}
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

