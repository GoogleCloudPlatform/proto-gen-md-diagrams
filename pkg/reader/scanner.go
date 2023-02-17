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
	"bufio"
	"os"
	"strings"
)

type Line struct {
	Syntax  string
	Token   string
	Comment string
}

// NewProtobufFileScanner is the constructor for ProtobufFileScanner
func NewProtobufFileScanner(file *os.File) Scanner {
	contents := strings.Join(readFileToArray(file), "\n")
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
	return NormalizeSpace(sw.scanner.Text())
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
	return newLine(sw.Text())
}

func newLine(in string) *Line {
	line := &Line{}
	if strings.HasPrefix(in, InlineCommentPrefix) {
		// Handle single comments
		line.Comment = strings.TrimSpace(in[strings.Index(in, InlineCommentPrefix)+len(InlineCommentPrefix):])
		line.Token = InlineCommentPrefix
	} else if strings.HasPrefix(in, MultiLineCommentInitiator) {
		// Handle Multiline Comments
		line.Comment = strings.TrimSpace(in[strings.Index(in, MultiLineCommentInitiator)+len(MultiLineCommentInitiator) : len(in)-len(MultilineCommentTerminator)])
		line.Token = MultiLineCommentInitiator
	} else if strings.Contains(in, Semicolon) {
		// Handle Syntax Stings
		line.Syntax = strings.TrimSpace(in[0:strings.Index(in, Semicolon)])
		line.Token = Semicolon
	} else if strings.Contains(in, OpenBrace) {
		// Handle Structure Strings
		line.Syntax = strings.TrimSpace(in[0:strings.Index(in, OpenBrace)])
		line.Token = OpenBrace
	} else if strings.Contains(in, ClosedBrace) {
		// Handle Inline Closed Structure Strings
		line.Syntax = strings.TrimSpace(in[0:strings.Index(in, ClosedBrace)])
		line.Token = ClosedBrace
	}
	// Add Inline Comments
	if !strings.HasPrefix(in, InlineCommentPrefix) && line.Token != MultiLineCommentInitiator && strings.Contains(in, InlineCommentPrefix) {
		line.Comment = Space + strings.TrimSpace(in[strings.Index(in, InlineCommentPrefix)+len(InlineCommentPrefix):])
	}
	line.Comment = TrimSpace(line.Comment)
	return line
}

// This non-exported function is responsible for reading the contents of
// a protocol buffer into a clean array of tokens.
func readFileToArray(file *os.File) []string {
	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	line := ""
	tokenReached := false

	for scanner.Scan() {
		rune := scanner.Text()
		if !strings.HasPrefix(line, MultiLineCommentInitiator) && (rune == Semicolon || rune == OpenBrace || rune == ClosedBrace) {
			lines = append(lines, line+rune, Space)
			tokenReached = true
			line = ""
		} else if strings.HasPrefix(line, InlineCommentPrefix) && rune == EndL {
			// Swap the comment with the tokenized line, accounting for: [{|;|}] //some comment
			if tokenReached {
				lines = append(lines, line, Space)
				// Swap first and last element
				pLine := lines[len(lines)-2]
				cLine := lines[len(lines)-1]
				lines[len(lines)-2] = cLine
				lines[len(lines)-1] = pLine
			} else {
				lines = append(lines, line, Space)
			}
			line = ""
			tokenReached = false
		} else if strings.HasPrefix(line, MultiLineCommentInitiator) && strings.HasSuffix(line, MultilineCommentTerminator) {
			lines = append(lines, line, Space)
			line = ""
		} else {
			if rune != EndL {
				if rune == Space {
					if len(line) > 0 {
						line += rune
					}
				} else {
					line += rune
				}
			} else {
				// Add a space to account for new lines in multiline comment
				if strings.HasPrefix(line, MultiLineCommentInitiator) {
					line += CommentNewLine
				}
				tokenReached = false
			}
		}
	}
	return lines
}
