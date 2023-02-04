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

// Package proto houses all logic for processing protocol buffers into
// an easy-to-use structure for creating diagrams output. The intention is to
// support a verity of diagram types such as Mermaid and Plant UML.
// Since Go does not support logical libraries, these are loaded via direction
// implementations. ToMermaid() and ToPlantUML(). Please note, since this is
// a Go implementation, the Plant UML diagrams are syntax only, and the Java
// diagram compiler is not currently used.
package proto

import "bufio"

// Validatable is a reference interface for the validator pattern
type Validatable interface {
	IsValid() bool
}

// Visitor is an interface used to determine if a line should be read,
// and if it should be, to read and interpret the line and subsequent lines
// as required.
type Visitor interface {
	CanVisit(in *Line) bool
	Visit(
			scanner Scanner,
			in *Line,
			namespace string) interface{}
}

// Scanner is an interface that SHOULD be a Go interface, but is only an
// implementation. Here, we can use the interface to wrap test cases
// with the same behavior of a bufio.Scanner
type Scanner interface {
	Scan() bool
	Text() string
	Split(splitFunction bufio.SplitFunc)
	Buffer(buf []byte, max int)
	Err() error
	Bytes() []byte
	ReadLine() *Line
}
