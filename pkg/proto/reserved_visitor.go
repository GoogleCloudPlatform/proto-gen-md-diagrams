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
	"strconv"
	"strings"
)

type ReservedVisitor struct {
}

func (rv *ReservedVisitor) CanVisit(line *Line) bool {
	return strings.HasPrefix(line.Syntax, PrefixReserved) && line.Token == Semicolon
}

func (rv *ReservedVisitor) Visit(_ Scanner, in *Line, _ string) interface{} {
	Log.Debug("Visiting Reserved")
	split := in.SplitSyntax()
	if len(split) == 2 {
		s, _ := strconv.ParseInt(split[1], 10, 64)
		return NewReserved(int32(s), int32(s))
	}
	if len(split) == 4 {
		s, _ := strconv.ParseInt(split[1], 10, 64)
		e, _ := strconv.ParseInt(split[3], 10, 64)
		return NewReserved(int32(s), int32(e))
	}
	return nil
}
