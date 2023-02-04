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

import (
	"fmt"
	"strings"
)

// PackageToMermaid formats a Package into Mermaid syntax
func PackageToMermaid(p *Package) string {
	out := fmt.Sprintf(`%%%%`+" Mermaid Diagram for package: %s\n", p.Name)

	for _, m := range p.Messages {
		out += MessageToMermaid(m)
	}

	for _, e := range p.Enums {
		out += EnumToMermaid(e)
	}

	for _, s := range p.Services {
		out += ServiceToMermaid(s)
	}

	return out
}

// EnumToMermaid formats an Enum into mermaid text.
func EnumToMermaid(e *Enum) string {
	out := fmt.Sprintf("%s\nclass %s{\n  <<enumeration>>\n", e.Comment.ToMermaid(), e.Name)
	for _, v := range e.Values {
		out += fmt.Sprintf("  %s\n", v.Value)
	}
	out += "}"
	return out
}

// MessageToMermaid formats a Message into mermaid text
func MessageToMermaid(m *Message) string {
	out := fmt.Sprintf("\n%s\nclass %s {\n", m.Comment.ToMermaid(), m.Name)
	for _, a := range m.Attributes {
		out += fmt.Sprintf("  %s\n", a.ToMermaid())
	}
	out += "}\n"

	// Handle Attribute Relationships
	for _, a := range m.Attributes {
		if len(a.Kind) == 1 {
			if !strings.Contains(Protobuf3Types, a.Kind[0]) {
				out += fmt.Sprintf("%s --> `%s`\n", m.Name, a.Kind[0])
			}
		} else if len(a.Kind) == 2 {
			if !strings.Contains(Protobuf3Types, strings.TrimSpace(a.Kind[1])) {
				out += fmt.Sprintf("%s .. `%s`\n", m.Name, a.Kind[1])
			}
		}
	}

	// Handle Message Relationships
	if m.HasMessages() {
		for _, msg := range m.Messages {
			out += fmt.Sprintf("%s --o `%s`\n", m.Name, msg.Name)
			out += MessageToMermaid(msg)
		}
	}

	// Handle Enumeration Relationships
	for _, e := range m.Enums {
		out += fmt.Sprintf("%s --o `%s`\n", m.Name, e.Name)
		out += EnumToMermaid(e)
	}

	return out
}

// FormatParametersForMermaid formats parameters for services
func FormatParametersForMermaid(in []*Parameter) string {
	out := ""
	for i := 0; i < len(in); i++ {
		p := in[i]
		if p.Stream {
			out += fmt.Sprintf("Stream~%s~", RemoveNameQualification(p.Type))
		} else {
			out += RemoveNameQualification(in[i].Type)
		}
		if i < len(in)-1 {
			out += ","
		}
	}
	return out
}

// FormatRelationships formats a service relationship
func FormatRelationships(name string, in []*Parameter) string {
	out := ""
	for _, i := range in {
		t := strings.TrimSpace(i.Type)
		if strings.HasSuffix(t, name) {
			t = RemoveNameQualification(t)
		}

		if i.Stream {
			out += fmt.Sprintf("%s --o `%s`\n", name, t)
		} else {
			out += fmt.Sprintf("%s --> `%s`\n", name, t)
		}
	}
	return out
}

// Formats a Service into mermaid text
func ServiceToMermaid(s *Service) string {
	relationships := ""

	out := fmt.Sprintf("class %s {\n  <<service>>\n", s.Name)
	for _, m := range s.Methods {
		out += fmt.Sprintf("  +%s(%s) %s\n",
			m.Name,
			FormatParametersForMermaid(m.InputParameters),
			FormatParametersForMermaid(m.ReturnParameters))

		relationships += FormatRelationships(s.Name, m.InputParameters)
		relationships += FormatRelationships(s.Name, m.ReturnParameters)
	}
	out += "}\n"
	out += relationships

	return out
}
