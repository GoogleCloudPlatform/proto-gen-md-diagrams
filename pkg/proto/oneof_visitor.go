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
	"strings"
)

// NewOneofVisitor creates a visitor for handling oneof blocks in protobuf messages.
func NewOneofVisitor() *OneofVisitor {
	visitors := make([]Visitor, 0)
	visitors = append(visitors, NewAttributeVisitor(), &CommentVisitor{})
	return &OneofVisitor{Visitors: visitors}
}

// OneofVisitor evaluates and marshals attributes declared inside a oneof block.
type OneofVisitor struct {
	Visitors []Visitor
}

// CanVisit determines if the line starts with 'oneof ' and ends with an open brace '{'
func (ov *OneofVisitor) CanVisit(in *Line) bool {
	return strings.HasPrefix(in.Syntax, "oneof ") && in.Token == OpenBrace
}

// Visit parses attributes within a oneof construct until the closing brace is reached.
func (ov *OneofVisitor) Visit(scanner Scanner, in *Line, namespace string) interface{} {
	Log.Debugf("Visiting Oneof: %v\n", in)

	values := in.SplitSyntax()
	out := NewOneof()
	if len(values) > 1 {
		out.Name = values[1]
	}
	out.Qualifier = Join(Period, namespace, out.Name)
	out.Comment = in.Comment

	var comment = Comment("")
	braceDepth := 1

	for scanner.Scan() {
		line := scanner.ReadLine()

		Log.Debugf("Scanning line in oneof: %s\n", line.Syntax)

		visited := false
		for _, visitor := range ov.Visitors {
			if visitor.CanVisit(line) {
				visited = true
				rt := visitor.Visit(
					scanner,
					line,
					namespace)
				switch t := rt.(type) {
				case *Attribute:
					if t.IsValid() {
						t.Comment = comment.AddSpace().Append(t.Comment).TrimSpace()
						t.Oneof = true
						t.OneofGroup = out.Name
						out.Attributes = append(out.Attributes, t)
						comment = comment.Clear()
					}
				case Comment:
					comment = comment.Append(t).AddSpace()
				}
				break
			}
		}
		if !visited {
			if line.Token == OpenBrace || strings.Contains(line.Syntax, OpenBrace) {
				braceDepth++
			} else if line.Token == CloseBrace || strings.Contains(line.Syntax, CloseBrace) {
				braceDepth--
				if braceDepth == 0 {
					break
				}
			}
		}
	}
	return out
}
