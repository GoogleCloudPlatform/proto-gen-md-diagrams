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

func TestGroupVisitor_CanVisit(t *testing.T) {
	type args struct {
		in *Line
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Can Visit Group",
			args: args{in: &Line{Syntax: "optional group Result = 1 {", Token: "{"}},
			want: true,
		},
		{
			name: "Can Visit Repeated Group",
			args: args{in: &Line{Syntax: "repeated group Detail = 2 {", Token: "{"}},
			want: true,
		},
		{
			name: "Can't Visit Standard Message",
			args: args{in: &Line{Syntax: "message User {", Token: "{"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gv := NewGroupVisitor()
			assert.Equal(t, tt.want, gv.CanVisit(tt.args.in))
		})
	}
}

func TestGroupVisitor_Visit(t *testing.T) {
	content := "  string url = 2;\n  optional int32 code = 3;\n}"
	testScanner := NewTestScanner(content)

	gv := NewGroupVisitor()
	line := &Line{Syntax: "optional group Result = 1 {", Token: "{"}
	res := gv.Visit(testScanner, line, "test.package")

	group, ok := res.(*Group)
	assert.True(t, ok)
	assert.NotNil(t, group)
	assert.Equal(t, "Result", group.Message.Name)
	assert.Equal(t, "result", group.Attribute.Name)
	assert.Equal(t, 1, group.Attribute.Ordinal)
	assert.True(t, group.Attribute.Optional)
	assert.Equal(t, 2, len(group.Message.Attributes))
}
