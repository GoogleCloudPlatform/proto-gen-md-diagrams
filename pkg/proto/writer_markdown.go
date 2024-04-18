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

type WriterConfig struct {
	visualize    bool
	pureMarkdown bool
}

func EnumToMarkdown(enum *Enum, wc *WriterConfig) (body string, diagram string) {
	enumTable := NewMarkdownTable()
	enumTable.AddHeader("Name", "Ordinal", "Description")
	for _, v := range enum.Values {
		if wc.pureMarkdown {
			enumTable.Insert(fmt.Sprintf("`%s`", v.Value), strconv.Itoa(v.Ordinal), v.Comment.ToMarkdownText(false))
		} else {
			enumTable.Insert(v.Value, strconv.Itoa(v.Ordinal), v.Comment.ToMarkdownText(false))
		}
	}

	// Convert to a string
	if wc.visualize {
		diagram = "\n" + ToMermaid(enum.Name, enum)
	}
	if wc.pureMarkdown {
		body = fmt.Sprintf("## Enum: %s\n\n%s\n\n%s\n\n%s\n\n", enum.Name, fmt.Sprintf(fqnPureMd, enum.Qualifier), enum.Comment.ToMarkdownText(true), enumTable.String())
	} else {
		body = fmt.Sprintf("## Enum: %s\n%s\n\n%s\n\n%s\n\n", enum.Name, fmt.Sprintf(fqn, enum.Qualifier), enum.Comment.ToMarkdownBlockQuote(), enumTable.String())
	}
	return body, diagram
}

func MessageToMarkdown(message *Message, wc *WriterConfig) (body string, diagram string) {
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
		if wc.pureMarkdown {
			attributeTable.Insert(fmt.Sprintf("`%s`", a.Name), strconv.Itoa(a.Ordinal), fmt.Sprintf("`%s`", strings.Join(a.Kind, Comma)), label, a.Comment.ToMarkdownText(false))
		} else {
			attributeTable.Insert(a.Name, strconv.Itoa(a.Ordinal), strings.Join(a.Kind, Comma), label, a.Comment.ToMarkdownText(false))
		}
	}

	if wc.visualize {
		diagram = "\n" + ToMermaid(message.Name, message)
	}

	if wc.pureMarkdown {
		body = fmt.Sprintf("## Message: %s\n\n%s\n\n%s\n\n%s\n\n", message.Name, fmt.Sprintf(fqnPureMd, message.Qualifier), message.Comment.ToMarkdownText(true), attributeTable.String())
	} else {
		body = fmt.Sprintf("## Message: %s\n%s\n\n%s\n\n%s\n\n", message.Name, fmt.Sprintf(fqn, message.Qualifier), message.Comment.ToMarkdownBlockQuote(), attributeTable.String())
	}
	for _, e := range message.Enums {
		eBody, eDiagram := EnumToMarkdown(e, wc)
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

func ServiceToMarkdown(s *Service, wc *WriterConfig) string {
	methodTable := NewMarkdownTable()
	methodTable.AddHeader("Method", "Parameter (In)", "Parameter (Out)", "Description")
	for _, m := range s.Methods {
		if wc.pureMarkdown {
			methodTable.Insert(fmt.Sprintf("`%s`", m.Name),
				fmt.Sprintf("`%s`", FormatServiceParameter(m.InputParameters)),
				fmt.Sprintf("`%s`", FormatServiceParameter(m.ReturnParameters)), m.Comment.ToMarkdownText(false))
		} else {
			methodTable.Insert(m.Name,
				FormatServiceParameter(m.InputParameters),
				FormatServiceParameter(m.ReturnParameters), m.Comment.ToMarkdownText(false))
		}
	}
	table := methodTable.String()
	if wc.visualize {
		table = ToMermaid(s.Name, s) + "\n\n" + table
	}

	if wc.pureMarkdown {
		return fmt.Sprintf("## Service: %s\n\n%s\n\n%s\n\n%s\n\n", s.Name, fmt.Sprintf(fqnPureMd, s.Qualifier), s.Comment.ToMarkdownText(true), table)
	}
	return fmt.Sprintf("## Service: %s\n%s\n\n%s\n\n%s\n\n", s.Name, fmt.Sprintf(fqn, s.Qualifier), s.Comment.ToMarkdownBlockQuote(), table)
}

func HandleEnums(enums []*Enum, wc *WriterConfig) (body string) {
	diagrams := ""
	if enums != nil {
		for _, e := range enums {
			eBody, eDiagram := EnumToMarkdown(e, wc)
			body += eBody
			diagrams += eDiagram
		}
	}
	return body + diagrams
}

func HandleMessages(messages []*Message, wc *WriterConfig) (body string) {
	diagrams := ""
	if messages != nil {
		for _, m := range messages {
			mBody, mDiagram := MessageToMarkdown(m, wc)
			body += mBody
			body += HandleMessages(m.Messages, wc)
			diagrams += mDiagram
		}
	}
	if wc.visualize {
		diagrams += "\n\n"
	}
	return diagrams + body
}

func PackageFormatImports(p *Package) (body string) {
	importTable := NewMarkdownTable()
	importTable.AddHeader("Import", "Description")
	for _, i := range p.Imports {
		importTable.Insert(i.Path, i.Comment.ToMarkdownText(false))
	}
	body = fmt.Sprintf("## Imports\n\n%s\n", importTable.String())
	return body
}

func PackageFormatOptions(p *Package) (body string) {
	optionTable := NewMarkdownTable()
	optionTable.AddHeader("Name", "Value", "Description")
	for _, o := range p.Options {
		optionTable.Insert(o.Name, o.Value, o.Comment.ToMarkdownText(false))
	}
	body = fmt.Sprintf("## Options\n\n%s\n", optionTable.String())
	return body
}

const fqn = "<div style=\"font-size: 12px; margin-top: -10px;\" class=\"fqn\">FQN: %s</div>"

const fqnPureMd = "**FQN**: %s"

const footer = `
<!-- Created by: Proto Diagram Tool -->
<!-- https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams -->`

func PackageToMarkDown(p *Package, wc *WriterConfig) string {
	out := ""
	if len(p.Services) > 0 {
		for _, s := range p.Services {
			out += ServiceToMarkdown(s, wc)
		}
	}
	out += HandleEnums(p.Enums, wc)
	out += HandleMessages(p.Messages, wc)
	if wc.pureMarkdown {
		out = fmt.Sprintf("# Package: %s\n\n%s\n\n%s\n\n%s\n\n%s\n%s\n", p.Name, p.Comment.ToMarkdownText(true), PackageFormatImports(p), PackageFormatOptions(p), out, footer)
	} else {
		out = fmt.Sprintf("# Package: %s\n\n%s\n\n%s\n\n%s\n\n%s\n%s\n", p.Name, p.Comment.ToMarkdownBlockQuote(), PackageFormatImports(p), PackageFormatOptions(p), out, footer)
	}
	return out
}
