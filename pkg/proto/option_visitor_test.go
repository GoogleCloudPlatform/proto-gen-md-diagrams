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

func TestOptionVisitor_CanVisit(t *testing.T) {
	type args struct {
		in *Line
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Can Visit",
			args: args{in: &Line{Syntax: "option java_package = \"com.google.test\"", Token: ";"}},
			want: true},
		{name: "Can't Visit",
			args: args{in: &Line{Syntax: "This is a comment", Token: "//"}},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ov := &OptionVisitor{}
			assert.Equalf(t, tt.want, ov.CanVisit(tt.args.in), "CanVisit(%v)", tt.args.in)
		})
	}
}

func TestOptionVisitor_Visit(t *testing.T) {

	testScanner := NewTestScanner("")

	type args struct {
		in0 Scanner
		in  *Line
		in2 string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "Visit", args: args{in0: testScanner, in: &Line{Syntax: "option java_package = \"com.github.rrmcguinness.proto.test.location\"", Token: ";"}, in2: "test"},
			want: &Option{NamedValue: &NamedValue{
				Name:  "java_package",
				Value: "com.github.rrmcguinness.proto.test.location",
			}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ov := &OptionVisitor{}
			assert.Equalf(t, tt.want, ov.Visit(tt.args.in0, tt.args.in, tt.args.in2), "Visit(%v, %v, %v)", tt.args.in0, tt.args.in, tt.args.in2)
		})
	}
}
