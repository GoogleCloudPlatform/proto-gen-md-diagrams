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
)

// PackageVisitor - Creates the package structure
type PackageVisitor struct {
}

func (pv *PackageVisitor) CanVisit(in *Line) bool {
	return strings.HasPrefix(in.Syntax, "package") && in.Token == Semicolon
}

func (pv *PackageVisitor) Visit(_ Scanner, in *Line, path string) interface{} {
	fValues := SplitSyntax(in.Syntax)
	return ProtobufFactory.NewPackage(path, fValues[1], in.Comment)
}

// ImportVisitor - Only applicable on Packages
type ImportVisitor struct {
}

func (iv *ImportVisitor) CanVisit(in *Line) bool {
	return strings.HasPrefix(in.Syntax, "import ") && in.Token == Semicolon
}

func (iv *ImportVisitor) Visit(_ Scanner, in *Line, _ string) interface{} {
	fValues := SplitSyntax(in.Syntax)
	return ProtobufFactory.NewImport(RemoveDoubleQuotes(RemoveSemicolon(fValues[1])), in.Comment)
}

// OptionVisitor - Only applicable on Packages
type OptionVisitor struct {
}

func (ov *OptionVisitor) CanVisit(in *Line) bool {
	return strings.HasPrefix(in.Syntax, "option ") && in.Token == Semicolon
}

func (ov *OptionVisitor) Visit(_ Scanner, in *Line, _ string) interface{} {
	fValues := SplitSyntax(in.Syntax)
	if len(fValues) == 4 {
		return ProtobufFactory.NewOption(
			fValues[1],
			strings.ReplaceAll(fValues[3], `"`, ""),
			in.Comment)
	}
	return ProtobufFactory.NewOption("Invalid", "", in.Comment)
}
