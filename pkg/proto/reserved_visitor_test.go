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

func TestReservedVisitor_CanVisit(t *testing.T) {
	type args struct {
		line *Line
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Can Visit", args: args{line: &Line{
			Syntax:  "reserved 10",
			Token:   ";",
			Comment: "Reserved 10",
		}}, want: true},
		{name: "Can Visit", args: args{line: &Line{
			Syntax:  "reserved 10 to 20",
			Token:   ";",
			Comment: "Reserved 10",
		}}, want: true},
		{name: "Can't Visit", args: args{line: &Line{
			Syntax: "Comment",
			Token:  "//",
		}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rv := &ReservedVisitor{}
			assert.Equalf(t, tt.want, rv.CanVisit(tt.args.line), "CanVisit(%v)", tt.args.line)
		})
	}
}

func TestReservedVisitor_Visit(t *testing.T) {
	type args struct {
		in0 Scanner
		in  *Line
		in2 string
	}
	testScanner := NewTestScanner(``)
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "Is Reserved", args: args{
			in0: testScanner,
			in: &Line{
				Syntax:  "reserved 10",
				Token:   ";",
				Comment: "Reserved 10",
			},
			in2: "test.Message",
		}, want: &Reserved{
			Start: 10,
			End:   10,
		}},
		{name: "Is Reserved", args: args{
			in0: testScanner,
			in: &Line{
				Syntax:  "reserved 10 to 20",
				Token:   ";",
				Comment: "Reserved 10 to 20",
			},
			in2: "test.Message",
		}, want: &Reserved{
			Start: 10,
			End:   20,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rv := &ReservedVisitor{}
			assert.Equalf(t, tt.want, rv.Visit(tt.args.in0, tt.args.in, tt.args.in2), "Visit(%v, %v, %v)", tt.args.in0, tt.args.in, tt.args.in2)
		})
	}
}
