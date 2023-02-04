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

func TestLine_SplitSyntax(t *testing.T) {
	type fields struct {
		Syntax  string
		Token   string
		Comment Comment
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{name: "Test Split Syntax", fields: fields{
			Syntax:  "message AddressType",
			Token:   ";",
			Comment: "Test",
		}, want: []string{"message", "AddressType"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Line{
				Syntax:  tt.fields.Syntax,
				Token:   tt.fields.Token,
				Comment: tt.fields.Comment,
			}
			assert.Equalf(t, tt.want, l.SplitSyntax(), "SplitSyntax()")
		})
	}
}

func TestNewLine(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want *Line
	}{
		{name: "Test Comment", args: args{in: "// Comment"}, want: &Line{Token: "//", Comment: "Comment"}},
		{name: "Test Multiline Comment", args: args{in: "/* Comment */"}, want: &Line{Token: "/*", Comment: "Comment"}},
		{name: "Test Open Brace", args: args{in: "message AddressType { // Comment"}, want: &Line{Token: "{", Syntax: "message AddressType", Comment: "Comment"}},
		{name: "Test Semicolon", args: args{in: "string name = 1; // Comment"}, want: &Line{Token: ";", Syntax: "string name = 1", Comment: "Comment"}},
		{name: "Test Close Brace", args: args{in: "} // Comment"}, want: &Line{Token: "}", Syntax: "", Comment: "Comment"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewLine(tt.args.in), "NewLine(%v)", tt.args.in)
		})
	}
}
