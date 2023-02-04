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

import "strings"

type ServiceVisitor struct {
	Visitors []Visitor
}

func NewServiceVisitor() *ServiceVisitor {
	visitors := make([]Visitor, 0)
	visitors = append(visitors, NewRpcVisitor(), &CommentVisitor{})
	return &ServiceVisitor{Visitors: visitors}
}

func (sv *ServiceVisitor) CanVisit(line *Line) bool {
	return strings.HasPrefix(line.Syntax, "service") && line.Token == OpenBrace
}

func (sv *ServiceVisitor) Visit(scanner Scanner, in *Line, namespace string) interface{} {
	Log.Debugf("Visiting Service: %v\n", in)

	values := in.SplitSyntax()
	out := NewService(namespace, values[1], in.Comment)

	comment := Comment("")

	for scanner.Scan() {
		line := scanner.ReadLine()
		Log.Debugf("Scanning line in service: %s", line.Syntax)
		if line.Token == CloseBrace {
			break
		}
		for _, visitor := range sv.Visitors {
			if visitor.CanVisit(line) {
				rt := visitor.Visit(scanner, line, Join(Period, namespace, out.Name))
				switch t := rt.(type) {
				case *Rpc:
					t.Comment = comment.AddSpace().Append(t.Comment).TrimSpace()
					out.AddRpc(t)
					comment = comment.Clear()
				case Comment:
					comment = comment.Append(t).AddSpace()
				}
			}
		}
	}
	return out
}
