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

type Qualified interface {
	Qualifier() string
	Name() string
	Comment() string
	SetComment(in string)
}

type ProtobufFactory interface {
	NewPackage(path string,
		name string,
		comment string) Package

	NewMessage(qualifier string,
		name string,
		comment string) Message

	NewReserved(start int32, end int32) Reserved

	NewService(path string,
		name string,
		comment string) Service

	NewAnnotation(name string, value string) Annotation

	NewAttribute(qualifier string,
		name string,
		comment string,
		isRepeated bool,
		isMap bool,
		ordinal int,
		kinds ...string,
	) Attribute

	NewEnum(
		qualifier string,
		name string,
		comment string,
		values ...EnumValue) Enum

	NewEnumValue(
		qualifier string,
		name string,
		comment string,
		ordinal int) EnumValue

	NewImport(path string, comment string) Import

	NewOption(name string, value string, comment string) Option

	NewRPC(qualifier string,
		name string,
		comment string) RPC

	NewRPCParameter(stream bool, kind string) RPCParameter
}

type Import interface {
	Path() string
	Comment() string
	SetComment(in string)
}

type Option interface {
	Name() string
	Value() string
	Comment() string
	SetComment(in string)
}

type Package interface {
	Qualified

	Options() []Option
	AddOption(name string, value string, comment string) Package

	Imports() []Import
	AddImport(path string, comment string) Package

	Messages() []Message
	AddMessage(message Message) Package

	Enums() []Enum
	AddEnum(enum Enum) Package

	Services() []Service
	AddService(service Service) Package

	GetGraph() Graph
}

type Message interface {
	Qualified

	Attributes() []Attribute
	AddAttribute(attribute Attribute) Message

	Messages() []Message
	AddMessage(message Message) Message

	Enums() []Enum
	AddEnum(enum Enum) Message

	Reserved() []Reserved
	AddReserved(start int32, end int32) Message

	GetGraph() Graph
}

type Annotation interface {
	Name() string
	Value() string
}

type Attribute interface {
	Qualified
	Validate() bool
	Repeated() bool
	Map() bool
	Kinds() []string
	Ordinal() int
	Annotations() []Annotation
	AddAnnotation(name string, value string) Attribute
}

type Service interface {
	Qualified
	RemoteProcedureCalls() []RPC
	AddRPC(rpc RPC) Service
	GetGraph() Graph
}

type Enum interface {
	Qualified
	Values() []EnumValue
	AddValue(value EnumValue) Enum
}

type EnumValue interface {
	Qualified
	Ordinal() int
}

type Reserved interface {
	Start() int32
	End() int32
}

type RPC interface {
	Qualified
	InputParameters() []RPCParameter
	AddInputParameter(stream bool, kind string) RPC

	ReturnParameters() []RPCParameter
	AddReturnParameter(stream bool, kind string) RPC

	RPCOptions() []RPCOption
	AddRPCOption(name string, comment string, body string) RPC
}

type RPCParameter interface {
	Stream() bool
	Kind() string
}

type RPCOption interface {
	Qualified
	Body() string
}

type Graph interface {
	AddVertex(v Vertex) error
	GetVertexByFQN(fqn string) (Vertex, error)
	GetVertex(vertexType VertexType, name string) (Vertex, error)
	Contains(v Vertex) bool
	VertexCount() int
	Vertices() []Vertex
}

type Vertex interface {
	FQN() string
	Type() VertexType
	Edges() []Vertex
	Properties() map[string]any
	EdgeCount() int
	AddEdge(other Vertex) error
}
