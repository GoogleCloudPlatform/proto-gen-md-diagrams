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

const (
	MarkdownPadding = 2
)

type MarkdownTable struct {
	Header        []string
	ColumnLengths []int
	Data          [][]string
}

func NewMarkdownTable() *MarkdownTable {
	return &MarkdownTable{Header: make([]string, 0), ColumnLengths: make([]int, 0), Data: make([][]string, 0)}
}

func (mt *MarkdownTable) EvaluateWidth(i int, d string) {
	dLen := len(d) + MarkdownPadding
	if len(mt.ColumnLengths) == i {
		mt.ColumnLengths = append(mt.ColumnLengths, dLen)
	} else if mt.ColumnLengths[i] < dLen {
		mt.ColumnLengths[i] = dLen
	}
}

func (mt *MarkdownTable) AddHeader(names ...string) {
	for i, d := range names {
		mt.EvaluateWidth(i, d)
		names[i] = Space + d
	}
	mt.Header = append(mt.Header, names...)
}

func (mt *MarkdownTable) Insert(data ...string) {
	for i, d := range data {
		mt.EvaluateWidth(i, d)
		// Pad
		data[i] = Space + d
	}
	mt.Data = append(mt.Data, data)
}

func ComputeFormat(length int, value string) string {
	out := value
	for i := 0; i < length-len(value); i++ {
		out += Space
	}
	out += Pipe
	return out
}

func DashLine(length int) string {
	return strings.Repeat(Hyphen, length) + Pipe
}

func (mt *MarkdownTable) String() string {
	// Write the Header
	out := Pipe
	for i, h := range mt.Header {
		out += ComputeFormat(mt.ColumnLengths[i], h)
	}
	// Write the Header Separator
	out += EndL + Pipe
	for i, _ := range mt.Header {
		out += DashLine(mt.ColumnLengths[i])
	}
	out += EndL
	// Write the Data
	for i := 0; i < len(mt.Data); i++ {
		out += Pipe
		for j := 0; j < len(mt.Data[i]); j++ {
			out += ComputeFormat(mt.ColumnLengths[j], mt.Data[i][j])
		}
		out += EndL
	}
	return out
}
