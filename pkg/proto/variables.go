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

import "errors"

// Effective Final Variables

// InvalidImport = used during import sequence of files.
var InvalidImport = errors.New("invalid import")

// Log is the Package Logger
var Log = &Logger{}

// SetDebug is used to enable the debug output, useful for troubleshooting.
func SetDebug(debug bool) {
	Log.debug = debug
}

var RegisteredVisitors []Visitor

// Initialize the Visitors
func init() {
	// Handle Comments
	RegisteredVisitors = append(RegisteredVisitors,
		&CommentVisitor{},
		&PackageVisitor{},
		&ImportVisitor{},
		&OptionVisitor{},
		&MessageVisitor{},
		&ReservedVisitor{},
		NewEnumVisitor(),
		NewAttributeVisitor(),
		NewServiceVisitor())
}
