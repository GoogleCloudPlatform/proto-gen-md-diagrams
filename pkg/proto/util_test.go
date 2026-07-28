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
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatLine(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "Test 001", args: args{in: " Hello     World    "}, want: "Hello World"},
		{name: "Test 002", args: args{in: "Hello     World    "}, want: "Hello World"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FormatLine(tt.args.in), "FormatLine(%v)", tt.args.in)
		})
	}
}

func TestJoin(t *testing.T) {
	type args struct {
		joinCharacter string
		values        []string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test 001", args: args{joinCharacter: ",", values: []string{"hello", "world"}}, want: "hello,world"},
		{name: "Test 001", args: args{joinCharacter: ".", values: []string{"hello", "world"}}, want: "hello.world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Join(tt.args.joinCharacter, tt.args.values...), "Join(%v, %v)", tt.args.joinCharacter, tt.args.values)
		})
	}
}

func TestNormalizeName(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test 001", args: args{in: "ThisIsATest"}, want: "this_is_a_test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NormalizeName(tt.args.in), "NormalizeName(%v)", tt.args.in)
		})
	}
}

func TestParseOrdinal(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test 001", args: args{in: "001"}, want: 1},
		{name: "Test 001", args: args{in: "10"}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ParseOrdinal(tt.args.in), "ParseOrdinal(%v)", tt.args.in)
		})
	}
}

func TestReadFileToArray(t *testing.T) {
	type args struct {
		file *os.File
	}

	fil, err := os.Open("data/input_test_file.txt")

	if err != nil {
		assert.Fail(t, "failed to open file", err)
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "Test 001", args: args{file: fil}, want: []string{"// this", "// is", "// for", "// testing"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ReadFileToArray(tt.args.file), "ReadFileToArray(%v)", tt.args.file)
		})
	}
}

func TestRemoveDoubleQuotes(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test 001", args: args{in: "\"test\""}, want: "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RemoveDoubleQuotes(tt.args.in), "RemoveDoubleQuotes(%v)", tt.args.in)
		})
	}
}

func TestRemoveNameQualification(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test 001", args: args{in: "com.google.Name"}, want: "Name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RemoveNameQualification(tt.args.in), "RemoveNameQualification(%v)", tt.args.in)
		})
	}
}

func TestRemoveSemicolon(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test 001", args: args{in: "test;"}, want: "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RemoveSemicolon(tt.args.in), "RemoveSemicolon(%v)", tt.args.in)
		})
	}
}

func TestReadRunesToArray(t *testing.T) {
	input := `message CancelDetailsExt {
/* XY id /
string orderId = 1;
/ Order cancellation code /
int32 code = 2;
/*

comment
description more
again more description
/
string percentage = 3;
/ Ticket id of ticket for cancellation /
string ticketId = 4;
/ Signature of ticket for cancellation */
string signature = 5;
}`

	expected := []string{
		"message CancelDetailsExt {",
		"/* XY id /",
		"string orderId = 1;",
		"/ Order cancellation code /",
		"int32 code = 2;",
		"/*:~::~:comment:~:description more:~:again more description:~:/",
		"string percentage = 3;",
		"/ Ticket id of ticket for cancellation /",
		"string ticketId = 4;",
		"/ Signature of ticket for cancellation */",
		"string signature = 5;",
		"}",
	}

	result := ReadRunesToArray(strings.NewReader(input))
	assert.Equal(t, expected, result)
}

func TestReadRunesToArray_CommentWithSemicolon(t *testing.T) {
	input := `syntax = "proto3";

package md.test;

// Example for issue
message Example {
  // This comment is ok
  int32 a = 1;
  // This comment have a semicolon: text;text
  int32 b = 2;
}`

	expected := []string{
		"syntax = \"proto3\";",
		"package md.test;",
		"// Example for issue",
		"message Example {",
		"// This comment is ok",
		"int32 a = 1;",
		"// This comment have a semicolon: text;text",
		"int32 b = 2;",
		"}",
	}

	result := ReadRunesToArray(strings.NewReader(input))
	assert.Equal(t, expected, result)
}
