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

package pb

import (
	"log"

	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/api"
)

// option an option implementation
type option struct {
	name    string
	value   string
	comment string
}

func (o *option) Name() string {
	return o.name
}

func (o *option) Value() string {
	return o.value
}

func (o *option) Comment() string {
	return o.comment
}

func (o *option) SetComment(in string) {
	o.comment = in
}

// _import an import implementation
type _import struct {
	path    string
	comment string
}

func (i *_import) Path() string {
	return i.path
}

func (i *_import) Comment() string {
	return i.comment
}

func (i *_import) SetComment(in string) {
	i.comment = in
}

// _package - the package implementation
type _package struct {
	api.Qualified
	options  []api.Option
	imports  []api.Import
	messages []api.Message
	enums    []api.Enum
	services []api.Service
	graph    api.Graph
}

func addVertex(pkg api.Package, vertexType api.VertexType, qualified api.Qualified) {
	err := pkg.GetGraph().AddVertex(api.NewVertex(qualified.Name(), vertexType))
	if err != nil {
		log.Default().Printf("failed to add vertex to package: %v", err)
	}
}

func (p *_package) Options() []api.Option {
	return p.options
}

func (p *_package) AddOption(name string, value string, comment string) api.Package {
	p.options = append(p.options, &option{name: name, value: value, comment: comment})
	return p
}

func (p *_package) Imports() []api.Import {
	return p.imports
}

func (p *_package) AddImport(path string, comment string) api.Package {
	p.imports = append(p.imports, &_import{path: path, comment: comment})
	return p
}

func (p *_package) Messages() []api.Message {
	return p.messages
}

func (p *_package) AddMessage(message api.Message) api.Package {
	addVertex(p, api.MESSAGE, message)
	p.messages = append(p.messages, message)
	return p
}

func (p *_package) Enums() []api.Enum {
	return p.enums
}

func (p *_package) AddEnum(enum api.Enum) api.Package {
	addVertex(p, api.ENUM, enum)
	p.enums = append(p.enums, enum)
	return p
}

func (p *_package) Services() []api.Service {
	return p.services
}

func (p *_package) AddService(service api.Service) api.Package {
	addVertex(p, api.SERVICE, service)
	p.services = append(p.services, service)
	return p
}

func (p _package) GetGraph() api.Graph {
	return p.graph
}
