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

// ExtendVisitor consumes proto2 extend blocks cleanly without syntax errors.
type ExtendVisitor struct {
}

func NewExtendVisitor() *ExtendVisitor {
	return &ExtendVisitor{}
}

func (ev *ExtendVisitor) CanVisit(in *Line) bool {
	return strings.HasPrefix(in.Syntax, "extend ") && (in.Token == OpenBrace || strings.Contains(in.Syntax, OpenBrace))
}

func (ev *ExtendVisitor) Visit(scanner Scanner, in *Line, namespace string) interface{} {
	Log.Debugf("Visiting Extend block: %v\n", in)
	braceDepth := 1
	for scanner.Scan() {
		line := scanner.ReadLine()
		if line.Token == OpenBrace || strings.Contains(line.Syntax, OpenBrace) {
			braceDepth++
		} else if line.Token == CloseBrace || strings.Contains(line.Syntax, CloseBrace) {
			braceDepth--
			if braceDepth == 0 {
				break
			}
		}
	}
	return nil
}
