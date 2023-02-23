/*
 * Copyright 2023 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"errors"
)

var VertexNotFound = errors.New("vertex not found")
var VertexExists = errors.New("vertex exists")
var EdgeExists = errors.New("edge exists")

func FQN(vertexType VertexType, name string) string {
	return string(vertexType) + ":" + name
}

func NewGraph() Graph {
	return &BaseGraph{vertices: make([]Vertex, 0)}
}

type BaseGraph struct {
	vertices []Vertex
}

func (g *BaseGraph) Vertices() []Vertex {
	return g.vertices
}

func (g *BaseGraph) AddVertex(v Vertex) error {
	if !g.Contains(v) {
		g.vertices = append(g.vertices, v)
		return nil
	}
	return VertexExists
}

func (g *BaseGraph) GetVertexByFQN(fqn string) (Vertex, error) {
	for _, n := range g.vertices {
		if n.FQN() == fqn {
			return n, nil
		}
	}
	return nil, VertexNotFound
}

func (g *BaseGraph) GetVertex(vertexType VertexType, name string) (Vertex, error) {
	return g.GetVertexByFQN(FQN(vertexType, name))
}

func (g *BaseGraph) Contains(v Vertex) bool {
	for _, n := range g.vertices {
		if n.FQN() == v.FQN() {
			return true
		}
	}
	return false
}

func (g *BaseGraph) VertexCount() int {
	return len(g.vertices)
}

func NewVertex(name string, typ VertexType, edges ...Vertex) Vertex {
	return &BaseVertex{
		name:  name,
		typ:   typ,
		edges: edges}
}

type BaseVertex struct {
	name       string
	typ        VertexType
	edges      []Vertex
	properties map[string]any
}

func (v *BaseVertex) FQN() string {
	return FQN(v.typ, v.name)
}

func (v *BaseVertex) Type() VertexType {
	return v.typ
}

func (v *BaseVertex) Edges() []Vertex {
	return v.edges
}

func (v *BaseVertex) Properties() map[string]any {
	return v.properties
}

func (v *BaseVertex) EdgeCount() int {
	return len(v.edges)
}

func (v *BaseVertex) AddEdge(other Vertex) error {
	if !v.IsRelated(other) {
		v.edges = append(v.edges, other)
		return nil
	}
	return EdgeExists
}

func (v *BaseVertex) IsRelated(other Vertex) bool {
	out := false
	for _, e := range v.edges {
		if e.FQN() == other.FQN() {
			out = true
			break
		}
	}
	return out
}
