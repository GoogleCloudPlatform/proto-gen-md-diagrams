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

package reader

import (
	"strings"

	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/api"
	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/logging"
)

// MessageVisitor is used for interpreting message text
type MessageVisitor struct {
	Log *logging.Logger
}

func NewMessageVisitor(debug bool) *MessageVisitor {
	return &MessageVisitor{Log: logging.NewLogger(debug, "message visitor")}
}

// CanVisit visits if the line starts with 'message' and ends with an open brace '{'
func (mv *MessageVisitor) CanVisit(in *Line) bool {
	return strings.HasPrefix(in.Syntax, "message ") && in.Token == OpenBrace
}

// Visit evaluates the current line and parses the message until the closed brace
// is evaluated.
func (mv *MessageVisitor) Visit(scanner Scanner, in *Line, namespace string) interface{} {
	mv.Log.Debugf("Visiting Message: %v\n", in)

	values := SplitSyntax(in.Syntax)
	out := ProtobufFactory.NewMessage(
		Join(Period, namespace, values[1]),
		values[1],
		in.Comment)

	var comment = ""

	for scanner.Scan() {
		line := scanner.ReadLine()

		mv.Log.Debugf("Current Line: `%s`\n", line)

		if strings.HasSuffix(line.Token, ClosedBrace) {
			break
		}
		for _, visitor := range RegisteredVisitors {
			if visitor.CanVisit(line) {
				rt := visitor.Visit(
					scanner,
					line,
					Join(Period, namespace, out.Name()))
				switch t := rt.(type) {
				case api.Message:
					t.SetComment(Join(Space, comment, t.Comment()))
					out.AddMessage(t)
					comment = ""
				case api.Enum:
					t.SetComment(Join(Space, comment, t.Comment()))
					out.AddEnum(t)
					comment = ""
				case api.Attribute:
					if t.Validate() {
						t.SetComment(Join(Space, comment, t.Comment()))
						out.AddAttribute(t)
						comment = ""
					}
				case api.Reserved:
					out.AddReserved(t.Start(), t.End())
				case string:
					comment = Join(Space, comment, t)
				}
			}
		}
	}
	return out
}

// NewAttributeVisitor - Constructor for the AttributeVisitor
func NewAttributeVisitor() *AttributeVisitor {
	return &AttributeVisitor{}
}

// AttributeVisitor implementation for attributes
type AttributeVisitor struct {
	Log logging.Logger
}

// CanVisit - Determines if the line is an attribute, it doesn't end in a brace,
// it's a map, repeated, or can effectively be split
func (av *AttributeVisitor) CanVisit(in *Line) bool {
	return (!strings.HasSuffix(in.Syntax, OpenBrace) || !strings.HasSuffix(in.Syntax, ClosedBrace)) &&
		strings.HasPrefix(in.Syntax, "repeated") ||
		strings.HasPrefix(in.Syntax, "map") || len(SplitSyntax(in.Syntax)) >= 4
}

// HandleRepeated marshals the attribute into a repeated representation, e.g. List.
func HandleRepeated(qualifier string,
	comment string,
	split []string) api.Attribute {
	return ProtobufFactory.NewAttribute(qualifier, split[2], comment, true, false, ParseOrdinal(split[4]), split[1])
}

// HandleMap marshals the attribute into a Map type by using multiple types for key and value.
func HandleMap(qualifier string,
	comment string,
	split []string) api.Attribute {

	mapValue := Join(Space, split[0], split[1])
	innerTypes := mapValue[strings.Index(mapValue, OpenMap)+len(OpenMap) : strings.Index(mapValue, ClosedMap)]
	splitTypes := strings.Split(innerTypes, Comma)

	return ProtobufFactory.NewAttribute(qualifier, split[2], comment, false, true,
		ParseOrdinal(split[4]), splitTypes...)

}

// HandleDefaultAttribute marshals a standard attribute type.
func HandleDefaultAttribute(qualifier string,
	comment string,
	split []string) api.Attribute {
	if len(split) >= 3 {
		return ProtobufFactory.NewAttribute(qualifier, split[1], comment, false, false, ParseOrdinal(split[3]), split[0])
	}
	return nil
}

// Visit is used for marshalling an attribute into a struct.
func (av *AttributeVisitor) Visit(_ Scanner, in *Line, namespace string) interface{} {
	av.Log.Debug("Visiting Attribute")

	split := SplitSyntax(in.Syntax)
	var comment = ""
	var out api.Attribute

	if strings.HasPrefix(in.Syntax, PrefixReserved) {
		av.Log.Debug("\t processing reserved attribute")
		comment += in.Comment
	} else if strings.HasPrefix(in.Syntax, PrefixRepeated) {
		out = HandleRepeated(namespace, in.Comment, split)
	} else if strings.HasPrefix(in.Syntax, PrefixMap) {
		out = HandleMap(namespace, in.Comment, split)
	} else {
		out = HandleDefaultAttribute(namespace, in.Comment, split)
	}

	if out != nil {
		var annotations = ParseAnnotations(in.Syntax)
		for _, a := range annotations {
			out.AddAnnotation(a.Name(), a.Value())
		}
	}

	return out
}

// ParseAnnotations is used for reading the annotation line and marshalling it into
// the annotation structure.
func ParseAnnotations(in string) []api.Annotation {
	out := make([]api.Annotation, 0)
	if strings.Contains(in, OpenBracket) && strings.Contains(in, ClosedBracket) {
		annotationString := in[strings.Index(in, OpenBracket)+1 : strings.Index(in, ClosedBracket)]
		split := strings.Split(strings.ReplaceAll(annotationString, SingleQuote, Empty), Space)
		out = append(out, ProtobufFactory.NewAnnotation(split[0], split[2]))
	}
	return out
}
