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
	"bufio"
	"os"
	"regexp"
	"strings"
)

var SpaceRemover *regexp.Regexp

func init() {
	SpaceRemover = regexp.MustCompile(SpaceRemovalRegex)
}

// NewProtobufFileScanner is the constructor for ProtobufFileScanner
func NewProtobufFileScanner(file *os.File) Scanner {
	contents := strings.Join(ReadFileToArray(file), "\n")
	scanner := bufio.NewScanner(strings.NewReader(contents))
	scanner.Split(bufio.ScanLines)
	return &ProtobufFileScanner{scanner: scanner}
}

// ProtobufFileScanner is a specialized scanner for reading protobuf 3 files.
type ProtobufFileScanner struct {
	scanner *bufio.Scanner
}

// Scan is a delegate method to the underline scanner
func (sw ProtobufFileScanner) Scan() bool {
	return sw.scanner.Scan()
}

// Text is a specialization of the Text function, ensuring the line read
// is ready for processing.
func (sw ProtobufFileScanner) Text() string {
	return FormatLine(sw.scanner.Text())
}

// Split is a delegate method to the underline scanner
func (sw ProtobufFileScanner) Split(splitFunction bufio.SplitFunc) {
	sw.scanner.Split(splitFunction)
}

// Buffer is a delegate method to the underline scanner
func (sw ProtobufFileScanner) Buffer(buf []byte, max int) {
	sw.scanner.Buffer(buf, max)
}

// Err is a delegate method to the underline scanner
func (sw ProtobufFileScanner) Err() error {
	return sw.scanner.Err()
}

// Bytes is a delegate method to the underline scanner
func (sw ProtobufFileScanner) Bytes() []byte {
	return sw.scanner.Bytes()
}

// ReadLine is an addition to the buffered reader responsible for interpreting
// the line of the protobuf for the AST.
func (sw ProtobufFileScanner) ReadLine() *Line {
	return NewLine(sw.Text())
}
