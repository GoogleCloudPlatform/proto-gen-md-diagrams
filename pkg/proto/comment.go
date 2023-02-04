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
	"fmt"
	"strings"
)

// Comment is a string with additional methods
type Comment string

func (c Comment) ToMermaid() string {
	comments := strings.Split(string(c), CommentNewLine)
	out := ""
	for _, c := range comments {
		out += fmt.Sprintf("%%%% %s\n", c)

	}
	return out
}

func (c Comment) ToMarkdownText() string {
	comments := strings.Split(string(c), CommentNewLine)
	out := ""
	for _, c := range comments {
		out += fmt.Sprintf("%s ", c)
	}
	return out
}

func (c Comment) ToMarkdownBlockQuote() string {
	comments := strings.Split(string(c), CommentNewLine)
	out := "<div class=\"comment\">"
	for _, c := range comments {
		out += fmt.Sprintf("<span>%s</span><br/>", c)
	}
	out += "</div>"
	return out
}

// Append adds a comment to the end of an existing comment.
func (c Comment) Append(other Comment) Comment {
	c += Space + Comment(strings.TrimSpace(string(other)))
	return c
}

// AddSpace adds a space to the end of a Comment.
func (c Comment) AddSpace() Comment {
	c += Space
	return c
}

// TrimSpace removes any double space or padding spaces from the comment.
func (c Comment) TrimSpace() Comment {
	return Comment(FormatLine(strings.TrimSpace(string(c))))
}

// Clear truncates the comment
func (c Comment) Clear() Comment {
	return c[:0]
}
