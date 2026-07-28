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
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageVisitor_CanVisit(t *testing.T) {
	type args struct {
		in *Line
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Can Visit", args: args{in: &Line{
			Syntax:  "message Test",
			Token:   "{",
			Comment: "Test Message",
		}}, want: true},
		{name: "Can not Visit", args: args{in: &Line{
			Token:   "//",
			Comment: "Test Message",
		}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mv := &MessageVisitor{}
			assert.Equalf(t, tt.want, mv.CanVisit(tt.args.in), "CanVisit(%v)", tt.args.in)
		})
	}
}

func TestMessageVisitor_Visit(t *testing.T) {
	type args struct {
		scanner   Scanner
		in        *Line
		namespace string
	}
	testFile := `
  enum TestEnum {
    T1 = 0;
    T2 = 1;
  }
  string name = 1; // Name
  TestEnum type = 2; // Type
`
	scanner := NewTestScanner(testFile)

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "Message Scanner", args: args{
			scanner:   scanner,
			in:        NewLine("message Test { // Test Message"),
			namespace: "test",
		}, want: &Message{
			Qualified: &Qualified{
				Qualifier: "test.Test",
				Name:      "Test",
				Comment:   "Test Message",
			},
			Attributes: []*Attribute{
				{
					Qualified: &Qualified{
						Qualifier: "test.Test",
						Name:      "name",
						Comment:   "Name",
					},
					Repeated:    false,
					Map:         false,
					Kind:        []string{"string"},
					Ordinal:     1,
					Annotations: make([]*Annotation, 0),
				},
				{
					Qualified: &Qualified{
						Qualifier: "test.Test",
						Name:      "type",
						Comment:   "Type",
					},
					Repeated:    false,
					Map:         false,
					Kind:        []string{"TestEnum"},
					Ordinal:     2,
					Annotations: make([]*Annotation, 0),
				},
			},
			Messages: make([]*Message, 0),
			Enums: []*Enum{
				{
					Qualified: &Qualified{
						Qualifier: "test.Test.TestEnum",
						Name:      "TestEnum",
					},
					Values: []*EnumValue{
						{
							Namespace: "test.Test.TestEnum",
							Ordinal:   0,
							Value:     "T1",
						},
						{
							Namespace: "test.Test.TestEnum",
							Ordinal:   1,
							Value:     "T2",
						},
					},
				},
			},
			Reserved: make([]*Reserved, 0),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mv := &MessageVisitor{}
			assert.Equalf(t, tt.want, mv.Visit(tt.args.scanner, tt.args.in, tt.args.namespace), "Visit(%v, %v, %v)", tt.args.scanner, tt.args.in, tt.args.namespace)
		})
	}
}

func TestMessageToMermaid_CancelDetailsExt(t *testing.T) {
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

	lines := ReadRunesToArray(strings.NewReader(input))
	contents := strings.Join(lines, "\n")
	scanner := bufio.NewScanner(strings.NewReader(contents))
	scanner.Split(bufio.ScanLines)
	pScanner := &ProtobufFileScanner{scanner: scanner}

	assert.True(t, pScanner.Scan())
	in := pScanner.ReadLine()

	mv := &MessageVisitor{}
	msg := mv.Visit(pScanner, in, "test").(*Message)

	mermaid := MessageToMermaid(msg)
	expectedMermaid := `
%% 

class CancelDetailsExt {
  + string orderId
  + int32 code
  + string percentage
  + string ticketId
  + string signature
}
`
	assert.Equal(t, expectedMermaid, mermaid)
}

func TestMessageVisitor_CommentWithSemicolon(t *testing.T) {
	input := `message Example {
  // This comment is ok
  int32 a = 1;
  // This comment have a semicolon: text;text
  int32 b = 2;
}`

	lines := ReadRunesToArray(strings.NewReader(input))
	contents := strings.Join(lines, "\n")
	scanner := bufio.NewScanner(strings.NewReader(contents))
	scanner.Split(bufio.ScanLines)
	pScanner := &ProtobufFileScanner{scanner: scanner}

	assert.True(t, pScanner.Scan())
	in := pScanner.ReadLine()

	mv := &MessageVisitor{}
	msg := mv.Visit(pScanner, in, "md.test").(*Message)

	assert.Equal(t, "Example", msg.Name)
	assert.Equal(t, 2, len(msg.Attributes))
	assert.Equal(t, "a", msg.Attributes[0].Name)
	assert.Equal(t, "This comment is ok", string(msg.Attributes[0].Comment))
	assert.Equal(t, "b", msg.Attributes[1].Name)
	assert.Equal(t, "This comment have a semicolon: text;text", string(msg.Attributes[1].Comment))
}

func TestMessageVisitor_Issue11_MultipleEnums(t *testing.T) {
	input := `message MessageA {
  enum ENUM_A {
    ENUM_A_UNSPECIFIED = 0;
    ENUM_A_1 = 1;
    ENUM_A_2 = 2;
  }
  ENUM_A enum_a = 1;
  enum ENUM_B {
    ENUM_B_UNSPECIFIED = 0;
    ENUM_B_1 = 1;
    ENUM_B_2 = 2;
  }
  ENUM_B enum_b = 2;
}`

	lines := ReadRunesToArray(strings.NewReader(input))
	contents := strings.Join(lines, "\n")
	scanner := bufio.NewScanner(strings.NewReader(contents))
	scanner.Split(bufio.ScanLines)
	pScanner := &ProtobufFileScanner{scanner: scanner}

	assert.True(t, pScanner.Scan())
	in := pScanner.ReadLine()

	mv := &MessageVisitor{}
	msg := mv.Visit(pScanner, in, "foo").(*Message)

	mermaid := MessageToMermaid(msg)
	assert.NotContains(t, mermaid, "}MessageA")
	assert.Contains(t, mermaid, "}\nMessageA --o `ENUM_B`")
}

func TestMessageVisitor_Issue11_CustomOptionBlock(t *testing.T) {
	input := `message MessageC {
  option (foo.custom) = {
    boo: true
  };

  string a = 1;
}`

	lines := ReadRunesToArray(strings.NewReader(input))
	contents := strings.Join(lines, "\n")
	scanner := bufio.NewScanner(strings.NewReader(contents))
	scanner.Split(bufio.ScanLines)
	pScanner := &ProtobufFileScanner{scanner: scanner}

	assert.True(t, pScanner.Scan())
	in := pScanner.ReadLine()

	mv := &MessageVisitor{}
	msg := mv.Visit(pScanner, in, "foo").(*Message)

	assert.Equal(t, "MessageC", msg.Name)
	assert.Equal(t, 1, len(msg.Attributes))
	assert.Equal(t, "a", msg.Attributes[0].Name)
	assert.Equal(t, []string{"string"}, msg.Attributes[0].Kind)
}
