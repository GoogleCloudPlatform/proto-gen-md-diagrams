package mermaid

import (
	"fmt"
	"log"

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

func UnifiedDiagram(p []api.Package) string {
	g := api.NewGraph()
	for _, pkg := range p {
		for _, s := range pkg.GetGraph().Vertices() {
			err := g.AddVertex(s)
			if err != nil {
				log.Default().Printf("Duplicate entry %s", s.FQN())
			}
		}
	}
	return ""
}

func ClassDiagram(p api.Package) string {
	out := codifyRelationships(p.Name(), p.GetGraph().Vertices())
	return out
}
