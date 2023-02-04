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

// Line is a split line for syntax, token, and comment
type Line struct {
	Syntax  string
	Token   string
	Comment Comment
}

func NewLine(in string) *Line {
	line := &Line{}
	if strings.HasPrefix(in, InlineCommentPrefix) {
		// Handle single comments
		line.Comment = Comment(strings.TrimSpace(in[strings.Index(in, InlineCommentPrefix)+len(InlineCommentPrefix):]))
		line.Token = InlineCommentPrefix
	} else if strings.HasPrefix(in, MultiLineCommentInitiator) {
		// Handle Multiline Comments
		line.Comment = Comment(strings.TrimSpace(in[strings.Index(in, MultiLineCommentInitiator)+len(MultiLineCommentInitiator) : len(in)-len(MultilineCommentTerminator)]))
		line.Token = MultiLineCommentInitiator
	} else if strings.Contains(in, Semicolon) {
		// Handle Syntax Stings
		line.Syntax = strings.TrimSpace(in[0:strings.Index(in, Semicolon)])
		line.Token = Semicolon
	} else if strings.Contains(in, OpenBrace) {
		// Handle Structure Strings
		line.Syntax = strings.TrimSpace(in[0:strings.Index(in, OpenBrace)])
		line.Token = OpenBrace
	} else if strings.Contains(in, CloseBrace) {
		// Handle Inline Closed Structure Strings
		line.Syntax = strings.TrimSpace(in[0:strings.Index(in, CloseBrace)])
		line.Token = CloseBrace
	}
	// Add Inline Comments
	if !strings.HasPrefix(in, InlineCommentPrefix) && line.Token != MultiLineCommentInitiator && strings.Contains(in, InlineCommentPrefix) {
		line.Comment = Comment(Space + strings.TrimSpace(in[strings.Index(in, InlineCommentPrefix)+len(InlineCommentPrefix):]))
	}
	line.Comment = line.Comment.TrimSpace()
	return line
}

// SplitSyntax breaks the syntax line on Space the character
func (l *Line) SplitSyntax() []string {
	return strings.Split(l.Syntax, Space)
}
