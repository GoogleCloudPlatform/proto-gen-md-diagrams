/*
 * Copyright 2023 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package proto

// An Attribute is a component in the message structure.
type Attribute struct {
	*Qualified
	Repeated    bool
	Optional    bool
	Map         bool
	Kind        []string
	Ordinal     int
	Annotations []*Annotation
}

// IsValid implements the Validatable interface
func (a *Attribute) IsValid() bool {
	return len(a.Name) > 0 && a.Kind != nil && len(a.Kind) >= 1 && a.Ordinal >= 1
}

// ToMermaid implements a Mermaid Syntax per Attribute
func (a *Attribute) ToMermaid() string {
	if a.Repeated {
		return Join("", "+ List~", a.Kind[0], "~ ", a.Name)
	} else if a.Map {
		return Join("", "+ Map~", a.Kind[0], ", ", a.Kind[1], "~ ", a.Name)
	} else if a.Optional {
		return Join("", "+ Optional~", a.Kind[0], "~ ", a.Name)
	}
	return Join(Space, "+", a.Kind[0], a.Name)
}

// NewAttribute is the Attribute constructor
func NewAttribute(namespace string, comment Comment) *Attribute {
	return &Attribute{
		Qualified:   &Qualified{Qualifier: namespace, Comment: comment},
		Repeated:    false,
		Annotations: make([]*Annotation, 0)}
}
