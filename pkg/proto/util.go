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
	"strconv"
	"strings"
	"unicode"
)

func RemoveSemicolon(in string) string {
	return strings.ReplaceAll(in, Semicolon, Empty)
}

func RemoveDoubleQuotes(in string) string {
	return strings.ReplaceAll(in, DoubleQuote, Empty)
}

func ParseOrdinal(in string) int {
	i, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		Log.Debugf("Failed to parse %s for integer", in)
		return 0
	}
	return int(i)
}

func FormatLine(in string) string {
	return strings.TrimSpace(SpaceRemover.ReplaceAllString(in, " "))
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

func NormalizeName(in string) string {
	clean := ""
	for i, r := range in {
		if i > 0 && unicode.IsUpper(r) {
			clean += Space
		}
		clean += string(r)
	}
	return strings.ReplaceAll(strings.ToLower(clean), Space, "_")
}

func ReadFileToArray(file *os.File) []string {
	cleaner := regexp.MustCompile(`\s+|\n`)
	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	line := ""
	tokenReached := false

	for scanner.Scan() {
		rune := scanner.Text()
		if !strings.HasPrefix(line, MultiLineCommentInitiator) && (rune == Semicolon || rune == OpenBrace || rune == CloseBrace) {
			lines = append(lines, cleaner.ReplaceAllString(line+rune, Space))
			tokenReached = true
			line = ""
		} else if strings.HasPrefix(line, InlineCommentPrefix) && rune == EndL {
			// Swap the comment with the tokenized line, accounting for: [{|;|}] //some comment
			if tokenReached {
				lines = append(lines, cleaner.ReplaceAllString(strings.TrimSpace(line), Space))
				// Swap first and last element
				pLine := lines[len(lines)-2]
				cLine := lines[len(lines)-1]
				lines[len(lines)-2] = cLine
				lines[len(lines)-1] = pLine
			} else {
				lines = append(lines, cleaner.ReplaceAllString(strings.TrimSpace(line), Space))
			}
			line = ""
			tokenReached = false
		} else if strings.HasPrefix(line, MultiLineCommentInitiator) && strings.HasSuffix(line, MultilineCommentTerminator) {
			lines = append(lines, cleaner.ReplaceAllString(strings.TrimSpace(line), Space))
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
	if Log.debug {
		for i, l := range lines {
			Log.Debugf("%d. %s", i, l)
		}
	}
	return lines
}

// RemoveNameQualification formats a parameter into a single name, this is due
// to a limitation in Mermaid that DOES NOT support fully qualified names.
func RemoveNameQualification(in string) string {
	if strings.Contains(in, Period) {
		return in[strings.LastIndex(in, Period)+1:]
	} else {
		return in
	}
}
