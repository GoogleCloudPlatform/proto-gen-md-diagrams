package mermaid

import (
	"fmt"
	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/api"
)

func codifyRelationships(root string, vertices []api.Vertex) string {
	var out = ""
	for _, v := range vertices {
		if v.Type() == api.MESSAGE {
			out += codifyRelationships(v.FQN(), v.Edges())
		}
		out += fmt.Sprintf("%s --> %s\n", root, v.FQN())
	}
	return out
}

func ClassDiagram(p api.Package) string {
	out := codifyRelationships(p.Name(), p.GetGraph().Vertices())
	return out
}
