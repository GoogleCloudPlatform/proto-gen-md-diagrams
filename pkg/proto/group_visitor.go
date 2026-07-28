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

// Group represents a proto2 group construct containing both a nested message and a field reference.
type Group struct {
	Message   *Message
	Attribute *Attribute
}

type GroupVisitor struct {
}

func NewGroupVisitor() *GroupVisitor {
	return &GroupVisitor{}
}

func (gv *GroupVisitor) CanVisit(in *Line) bool {
	return (strings.HasPrefix(in.Syntax, "group ") ||
		strings.HasPrefix(in.Syntax, "optional group ") ||
		strings.HasPrefix(in.Syntax, "repeated group ") ||
		strings.HasPrefix(in.Syntax, "required group ")) &&
		(in.Token == OpenBrace || strings.Contains(in.Syntax, OpenBrace))
}

func (gv *GroupVisitor) Visit(scanner Scanner, in *Line, namespace string) interface{} {
	Log.Debugf("Visiting Group: %v\n", in)
	split := in.SplitSyntax()

	groupName := ""
	var ordinal int
	isRepeated := strings.HasPrefix(in.Syntax, "repeated")
	isOptional := strings.HasPrefix(in.Syntax, "optional")

	for idx, val := range split {
		if val == "group" && idx+1 < len(split) {
			groupName = split[idx+1]
		}
		if val == "=" && idx+1 < len(split) {
			ordinal = ParseOrdinal(split[idx+1])
		}
	}

	if groupName == "" {
		return nil
	}

	fieldName := strings.ToLower(groupName)

	msg := NewMessage()
	msg.Name = groupName
	msg.Qualifier = Join(Period, namespace, msg.Name)
	msg.Comment = in.Comment

	var comment = Comment("")
	braceDepth := 1

	for scanner.Scan() {
		line := scanner.ReadLine()
		visited := false
		for _, visitor := range RegisteredVisitors {
			if visitor.CanVisit(line) {
				visited = true
				rt := visitor.Visit(scanner, line, Join(Period, namespace, msg.Name))
				switch t := rt.(type) {
				case *Message:
					msg.Messages = append(msg.Messages, t)
				case *Enum:
					msg.Enums = append(msg.Enums, t)
				case *Attribute:
					if t.IsValid() {
						msg.Attributes = append(msg.Attributes, t)
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

	attr := NewAttribute(namespace, in.Comment)
	attr.Name = fieldName
	attr.Kind = append(attr.Kind, groupName)
	attr.Ordinal = ordinal
	attr.Repeated = isRepeated
	attr.Optional = isOptional

	return &Group{Message: msg, Attribute: attr}
}
