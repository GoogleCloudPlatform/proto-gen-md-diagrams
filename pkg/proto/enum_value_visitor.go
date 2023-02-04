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

// EnumValueVisitor is responsible evaluating and processing Protobuf Enumerations.
type EnumValueVisitor struct {
}

// CanVisit determines if the line is an enumeration.
func (evv EnumValueVisitor) CanVisit(in *Line) bool {
	a := in.SplitSyntax()
	return a != nil && len(a) == 3 && in.Token == Semicolon
}

// Visit marshals a line into an enumeration
func (evv EnumValueVisitor) Visit(_ Scanner, in *Line, namespace string) interface{} {
	a := in.SplitSyntax()
	return NewEnumValue(namespace, a[2], a[0], in.Comment)
}
