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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeFormat(t *testing.T) {
	type args struct {
		length int
		value  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Good Format", args: args{length: 10, value: "test"}, want: "test      |"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ComputeFormat(tt.args.length, tt.args.value), "ComputeFormat(%v, %v)", tt.args.length, tt.args.value)
		})
	}
}

func TestDashLine(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, DashLine(tt.args.length), "DashLine(%v)", tt.args.length)
		})
	}
}

func TestMarkdownTable_AddHeader(t *testing.T) {
	type fields struct {
		header        []string
		columnLengths []int
		data          [][]string
	}
	type args struct {
		names []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Add Header",
			args: args{names: []string{"test"}},
			fields: fields{
				header:        make([]string, 0),
				columnLengths: make([]int, 0),
				data:          make([][]string, 0),
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := &MarkdownTable{
				Header:        tt.fields.header,
				ColumnLengths: tt.fields.columnLengths,
				Data:          tt.fields.data,
			}
			mt.AddHeader(tt.args.names...)
			assert.Equal(t, 1, len(mt.Header))
		})
	}
}

func TestMarkdownTable_EvaluateWidth(t *testing.T) {
	type fields struct {
		header        []string
		columnLengths []int
		data          [][]string
	}
	type args struct {
		i int
		d string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Evaluate Width",
			fields: fields{
				header:        make([]string, 0),
				columnLengths: make([]int, 0),
				data:          make([][]string, 0),
			}, args: args{
			i: 0,
			d: "test",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := &MarkdownTable{
				Header:        tt.fields.header,
				ColumnLengths: tt.fields.columnLengths,
				Data:          tt.fields.data,
			}
			mt.EvaluateWidth(tt.args.i, tt.args.d)
			assert.Equal(t, 6, mt.ColumnLengths[0])
		})
	}
}

func TestMarkdownTable_Insert(t *testing.T) {
	type fields struct {
		header        []string
		columnLengths []int
		data          [][]string
	}
	type args struct {
		data []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Insert",
			fields: fields{
				header:        make([]string, 0),
				columnLengths: make([]int, 0),
				data:          make([][]string, 0),
			}, args: args{data: []string{"test1", "test2", "test3"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := &MarkdownTable{
				Header:        tt.fields.header,
				ColumnLengths: tt.fields.columnLengths,
				Data:          tt.fields.data,
			}
			mt.Insert(tt.args.data...)
			assert.Equal(t, 3, len(mt.Data[0]))
		})
	}
}

func TestMarkdownTable_String(t *testing.T) {
	type fields struct {
		header        []string
		columnLengths []int
		data          [][]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// The padding in this test is because it's not using
		// the Insert method.
		{
			name: "String Test",
			fields: fields{
				header:        []string{" c1"},
				columnLengths: []int{6},
				data:          [][]string{{" test"}},
			},
			want: "| c1   |\n|------|\n| test |\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := &MarkdownTable{
				Header:        tt.fields.header,
				ColumnLengths: tt.fields.columnLengths,
				Data:          tt.fields.data,
			}
			assert.Equalf(t, tt.want, mt.String(), "String()")
		})
	}
}

func TestNewMarkdownTable(t *testing.T) {
	tests := []struct {
		name string
		want *MarkdownTable
	}{
		{name: "New Table", want: &MarkdownTable{
			Header:        []string{},
			ColumnLengths: []int{},
			Data:          [][]string{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewMarkdownTable(), "NewMarkdownTable()")
		})
	}
}
