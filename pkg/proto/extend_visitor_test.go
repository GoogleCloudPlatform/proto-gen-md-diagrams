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

func TestExtendVisitor_CanVisit(t *testing.T) {
	type args struct {
		in *Line
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Can Visit Extend",
			args: args{in: &Line{Syntax: "extend google.protobuf.FieldOptions {", Token: "{"}},
			want: true,
		},
		{
			name: "Can't Visit Option",
			args: args{in: &Line{Syntax: "option java_package = \"com.test\"", Token: ";"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev := NewExtendVisitor()
			assert.Equal(t, tt.want, ev.CanVisit(tt.args.in))
		})
	}
}

func TestExtendVisitor_Visit(t *testing.T) {
	content := "  optional string my_option = 50000;\n}"
	testScanner := NewTestScanner(content)

	ev := NewExtendVisitor()
	line := &Line{Syntax: "extend google.protobuf.FieldOptions {", Token: "{"}
	res := ev.Visit(testScanner, line, "test.package")
	assert.Nil(t, res)
}
