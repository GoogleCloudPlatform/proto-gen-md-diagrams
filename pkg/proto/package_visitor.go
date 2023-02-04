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

var isDebug bool

type PackageVisitor struct {
}

func (pv *PackageVisitor) CanVisit(in *Line) bool {
	return strings.HasPrefix(in.Syntax, "package") && in.Token == Semicolon
}

func (pv *PackageVisitor) Visit(_ Scanner, in *Line, _ string) interface{} {
	fValues := in.SplitSyntax()
	return &Package{
		Path:     "",
		Name:     fValues[1],
		Comment:  in.Comment,
		Options:  make([]*Option, 0),
		Imports:  make([]*Import, 0),
		Messages: make([]*Message, 0),
		Enums:    make([]*Enum, 0),
		Services: make([]*Service, 0),
	}
}
