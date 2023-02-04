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

type ImportVisitor struct {
}

func (iv *ImportVisitor) CanVisit(in *Line) bool {
	return strings.HasPrefix(in.Syntax, "import ") && in.Token == Semicolon
}

func (iv *ImportVisitor) Visit(_ Scanner, in *Line, _ string) interface{} {
	Log.Debug("Visiting Import")
	fValues := in.SplitSyntax()
	return NewImport(RemoveDoubleQuotes(RemoveSemicolon(fValues[1])))
}
