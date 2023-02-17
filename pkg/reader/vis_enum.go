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

// NewEnumVisitor creates an EnumVisitor
func NewEnumVisitor(debug bool) *EnumVisitor {
	out := &EnumVisitor{Log: logging.NewLogger(debug, "enum visitor"), Visitors: make([]Visitor, 0)}
	out.Visitors = append(out.Visitors,
		&CommentVisitor{},
		&EnumValueVisitor{})
	return out
}

// EnumVisitor is responsible for evaluation and marshalling of an Enum entity.
type EnumVisitor struct {
	Log      *logging.Logger
	Visitors []Visitor
}

// CanVisit determines if the current line is an enumeration.
func (ev *EnumVisitor) CanVisit(in *Line) bool {
	return strings.HasPrefix(in.Syntax, "enum ") && in.Token == OpenBrace
}

// Visit marshals a line and subsequent lines of the enumeration until the terminator is found.
func (ev *EnumVisitor) Visit(scanner Scanner, in *Line, namespace string) interface{} {
	ev.Log.Debugf("Visiting Enum: %d registered visitors\n", len(ev.Visitors))
	fValues := SplitSyntax(in.Syntax)
	out := ProtobufFactory.NewEnum(Join(Period, namespace, fValues[1]), fValues[1], in.Comment)

	var comment = ""

	for scanner.Scan() {
		n := scanner.ReadLine()
		if strings.HasSuffix(n.Token, ClosedBrace) {
			break
		}
		for _, visitor := range ev.Visitors {
			if visitor.CanVisit(n) {
				rt := visitor.Visit(
					scanner,
					n,
					Join(Period, namespace, out.Name()))

				switch t := rt.(type) {
				case api.EnumValue:
					t.SetComment(Join(Space, comment, t.Comment()))
					out.AddValue(t)
					comment = ""
				case string:
					comment = Join(Space, comment, t)
				default:
					ev.Log.Infof("unable to parse enum value: %t", t)
				}
			}
		}
	}
	return out
}

// EnumValueVisitor is responsible evaluating and processing Protobuf Enumerations.
type EnumValueVisitor struct {
}

// CanVisit determines if the line is an enumeration.
func (evv EnumValueVisitor) CanVisit(in *Line) bool {
	a := SplitSyntax(in.Syntax)
	return a != nil && len(a) == 3 && in.Token == Semicolon
}

// Visit marshals a line into an enumeration
func (evv EnumValueVisitor) Visit(_ Scanner, in *Line, namespace string) interface{} {
	a := SplitSyntax(in.Syntax)
	return ProtobufFactory.NewEnumValue(namespace, a[0], in.Comment, ParseOrdinal(a[2]))
}
