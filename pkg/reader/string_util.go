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
	"regexp"
	"strconv"
	"strings"
)

var SpaceRemover *regexp.Regexp

func init() {
	SpaceRemover = regexp.MustCompile(SpaceRemovalRegex)
}

func Join(joinCharacter string, values ...string) string {
	out := ""
	count := len(values)
	for i := 0; i < count; i++ {
		if i < count-1 {
			out += values[i] + joinCharacter
		} else {
			out += values[i]
		}
	}
	return out
}

func TrimSpace(in string) string {
	return strings.TrimSpace(in)
}

func NormalizeSpace(in string) string {
	return strings.TrimSpace(SpaceRemover.ReplaceAllString(in, " "))
}

func SplitSyntax(in string) []string {
	return strings.Split(in, Space)
}

func SplitComment(in string) string {
	return strings.ReplaceAll(in, CommentNewLine, EndL)
}

func RemoveSemicolon(in string) string {
	return strings.ReplaceAll(in, Semicolon, Empty)
}

func RemoveDoubleQuotes(in string) string {
	return strings.ReplaceAll(in, DoubleQuote, Empty)
}

func ParseOrdinal(in string) int {
	i, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return 0
	}
	return int(i)
}
