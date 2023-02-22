package test

import (
	"testing"

	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestSimpleGraph(t *testing.T) {

	g := api.NewGraph()

	assert.NotNil(t, g)

	e1 := g.AddVertex(api.NewVertex("m1", api.MESSAGE, api.NewVertex("m1e1", api.ENUM)))
	assert.Nil(t, e1)

	assert.Equal(t, 1, g.VertexCount())
	m1, e2 := g.GetVertex(api.MESSAGE, "m1")

	assert.Nil(t, e2)
	assert.NotNil(t, m1)

	m2, e3 := g.GetVertex(api.ENUM, "m1")
	assert.Nil(t, m2)
	assert.NotNil(t, e3)

	assert.Equal(t, 1, m1.EdgeCount())
	assert.Equal(t, api.VertexNotFound, e3)

	e4 := g.AddVertex(api.NewVertex("m1", api.MESSAGE))
	assert.NotNil(t, e4)
	assert.Equal(t, api.VertexExists, e4)
}
