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

import "github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/api"

type message struct {
	api.Qualified
	attributes []api.Attribute
	messages   []api.Message
	enums      []api.Enum
	reserved   []api.Reserved
}

func (m *message) Attributes() []api.Attribute {
	return m.Attributes()
}

func (m *message) AddAttribute(attribute api.Attribute) api.Message {
	m.attributes = append(m.attributes, attribute)
	return m
}

func (m *message) Messages() []api.Message {
	return m.messages
}

func (m *message) AddMessage(message api.Message) api.Message {
	m.messages = append(m.messages, message)
	return m
}

func (m *message) Enums() []api.Enum {
	return m.enums
}

func (m *message) AddEnum(enum api.Enum) api.Message {
	m.enums = append(m.enums, enum)
	return m
}

func (m *message) Reserved() []api.Reserved {
	return m.reserved
}

func (m *message) AddReserved(start int32, end int32) api.Message {
	m.reserved = append(m.reserved, &reserved{start: start, end: end})
	return m
}

// reserved the default implementation for api.Reserved
type reserved struct {
	start int32
	end   int32
}

func (r *reserved) Start() int32 {
	return r.start
}

func (r *reserved) End() int32 {
	return r.end
}

// annotation implements api.Annotation
type annotation struct {
	name  string
	value string
}

func (a *annotation) Name() string {
	return a.name
}

func (a *annotation) Value() string {
	return a.value
}

// attribute is the implementation for api.Attribute
type attribute struct {
	api.Qualified
	repeated    bool
	isMap       bool
	kinds       []string
	ordinal     int
	annotations []api.Annotation
}

func (a *attribute) Validate() bool {
	return len(a.Name()) > 0 && a.Kinds() != nil && len(a.Kinds()) >= 1 && a.ordinal >= 1
}

func (a *attribute) Repeated() bool {
	return a.repeated
}

func (a *attribute) Map() bool {
	return a.isMap
}

func (a *attribute) Kinds() []string {
	return a.kinds
}

func (a *attribute) Annotations() []api.Annotation {
	return a.annotations
}

func (a *attribute) Ordinal() int {
	return a.ordinal
}

func (a *attribute) AddAnnotation(name string, value string) api.Attribute {
	a.annotations = append(a.annotations, &annotation{name: name, value: value})
	return a
}
