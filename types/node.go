package types

import (
	"github.com/gabesullice/phargo/schema"
)

var (
	nodeSchema = schema.Schema{
		Type: "object",
		Id:   "https://www.phargo.io/schemas/0.0.0/node.json",
		Properties: map[string]schema.Schema{
			"id":      {Type: []string{"string"}},
			"created": {Type: []string{"string"}},
			"updated": {Type: []string{"string"}},
		},
		RequiredProperties: []string{
			"id",
			"created",
			"updated",
		},
	}
)

type Node struct{}

func (n Node) Schema() schema.Schema {
	return nodeSchema
}
