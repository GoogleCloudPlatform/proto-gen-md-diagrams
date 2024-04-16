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
	"sort"
	"strconv"
	"strings"
)

const (
	mermaidClassDiagramTemplate = "### %s Diagram\n\n```mermaid\nclassDiagram\ndirection LR\n%s\n```"
)

func ToMermaid(title string, rt interface{}) string {
	out := ""
	switch t := rt.(type) {
	case *Package:
		out += PackageToMermaid(t)
	case *Enum:
		out += EnumToMermaid(t)
	case *Message:
		out += MessageToMermaid(t)
	case *Service:
		out += ServiceToMermaid(t)
	}
	return fmt.Sprintf(mermaidClassDiagramTemplate, title, out)
}

func EnumToMarkdown(enum *Enum, visualize bool) (body string, diagram string) {
	enumTable := NewMarkdownTable()
	enumTable.AddHeader("Name", "Ordinal", "Description")
	for _, v := range enum.Values {
		enumTable.Insert(v.Value, strconv.Itoa(v.Ordinal), v.Comment.ToMarkdownText())
	}

	// Convert to a string
	if visualize {
		diagram = "\n" + ToMermaid(enum.Name, enum)
	}
	body = fmt.Sprintf("## Enum: %s\n%s\n\n%s\n\n%s\n\n", enum.Name, fmt.Sprintf(fqn, enum.Qualifier), enum.Comment.ToMarkdownBlockQuote(), enumTable.String())
	return body, diagram
}

func MessageToMarkdown(message *Message, visualize bool) (body string, diagram string) {
	attributeTable := NewMarkdownTable()
	attributeTable.AddHeader("Field", "Ordinal", "Type", "Label", "Description")

	sort.Slice(message.Attributes, func(i, j int) bool {
		return message.Attributes[i].Ordinal < message.Attributes[j].Ordinal
	})

	for _, a := range message.Attributes {
		label := ""
		if a.Map {
			label = "Map"
		} else if a.Repeated {
			label = "Repeated"
		} else if a.Optional {
			label = "Optional"
		}
		attributeTable.Insert(a.Name, strconv.Itoa(a.Ordinal), strings.Join(a.Kind, Comma), label, a.Comment.ToMarkdownText())
	}

	if visualize {
		diagram = "\n" + ToMermaid(message.Name, message)
	}

	body = fmt.Sprintf("## Message: %s\n%s\n\n%s\n\n%s\n\n", message.Name, fmt.Sprintf(fqn, message.Qualifier), message.Comment.ToMarkdownBlockQuote(), attributeTable.String())

	for _, e := range message.Enums {
		eBody, eDiagram := EnumToMarkdown(e, visualize)
		body += eBody
		diagram += eDiagram
	}
	return body, diagram
}

func FormatServiceParameter(parameters []*Parameter) string {
	out := ""
	for _, p := range parameters {
		if p.Stream {
			out += fmt.Sprintf("Stream\\<%s\\>", RemoveNameQualification(p.Type))
		} else {
			out += fmt.Sprintf("%s", RemoveNameQualification(p.Type))
		}
	}
	return out
}

func ServiceToMarkdown(s *Service, visualize bool) string {
	methodTable := NewMarkdownTable()
	methodTable.AddHeader("Method", "Parameter (In)", "Parameter (Out)", "Description")
	for _, m := range s.Methods {
		methodTable.Insert(m.Name,
			FormatServiceParameter(m.InputParameters),
			FormatServiceParameter(m.ReturnParameters), m.Comment.ToMarkdownText())
	}
	table := methodTable.String()
	if visualize {
		table = ToMermaid(s.Name, s) + "\n\n" + table
	}

	return fmt.Sprintf("## Service: %s\n%s\n\n%s\n\n%s\n\n", s.Name, fmt.Sprintf(fqn, s.Qualifier), s.Comment.ToMarkdownBlockQuote(), table)
}

func HandleEnums(enums []*Enum, visualize bool) (body string) {
	diagrams := ""
	if enums != nil {
		for _, e := range enums {
			eBody, eDiagram := EnumToMarkdown(e, visualize)
			body += eBody
			diagrams += eDiagram
		}
	}
	return body + diagrams
}

func HandleMessages(messages []*Message, visualize bool) (body string) {
	diagrams := ""
	if messages != nil {
		for _, m := range messages {
			mBody, mDiagram := MessageToMarkdown(m, visualize)
			body += mBody
			body += HandleMessages(m.Messages, false)
			diagrams += mDiagram
		}
	}
	if visualize {
		diagrams += "\n\n"
	}
	return diagrams + body
}

func PackageFormatImports(p *Package) (body string) {
	importTable := NewMarkdownTable()
	importTable.AddHeader("Import", "Description")
	for _, i := range p.Imports {
		importTable.Insert(i.Path, i.Comment.ToMarkdownText())
	}
	body = fmt.Sprintf("## Imports\n\n%s\n", importTable.String())
	return body
}

func PackageFormatOptions(p *Package) (body string) {
	optionTable := NewMarkdownTable()
	optionTable.AddHeader("Name", "Value", "Description")
	for _, o := range p.Options {
		optionTable.Insert(o.Name, o.Value, o.Comment.ToMarkdownText())
	}
	body = fmt.Sprintf("## Options\n\n%s\n", optionTable.String())
	return body
}

const fqn = "<div style=\"font-size: 12px; margin-top: -10px;\" class=\"fqn\">FQN: %s</div>"

const footer = `
<!-- Created by: Proto Diagram Tool -->
<!-- https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams -->`

func PackageToMarkDown(p *Package, visualize bool) string {
	out := ""
	if len(p.Services) > 0 {
		for _, s := range p.Services {
			out += ServiceToMarkdown(s, visualize)
		}
	}
	out += HandleEnums(p.Enums, visualize)
	out += HandleMessages(p.Messages, visualize)
	out = fmt.Sprintf("# Package: %s\n\n%s\n\n%s\n\n%s\n\n%s\n%s\n", p.Name, p.Comment.ToMarkdownBlockQuote(), PackageFormatImports(p), PackageFormatOptions(p), out, footer)
	return out
}
