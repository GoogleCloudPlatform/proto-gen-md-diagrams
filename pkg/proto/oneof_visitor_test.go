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

func TestOneofVisitor_CanVisit(t *testing.T) {
	type args struct {
		in *Line
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Can Visit oneof",
			args: args{in: &Line{
				Syntax:  "oneof test_oneof",
				Token:   "{",
				Comment: "Test Oneof",
			}},
			want: true,
		},
		{
			name: "Can not Visit message",
			args: args{in: &Line{
				Syntax:  "message Test",
				Token:   "{",
				Comment: "Test Message",
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ov := NewOneofVisitor()
			assert.Equalf(t, tt.want, ov.CanVisit(tt.args.in), "CanVisit(%v)", tt.args.in)
		})
	}
}

func TestOneofVisitor_Visit(t *testing.T) {
	testFile := `
  string name = 1; // Oneof field 1
  int32 id = 2; // Oneof field 2
}
`
	scanner := NewTestScanner(testFile)
	ov := NewOneofVisitor()
	in := NewLine("oneof test_oneof {")

	got := ov.Visit(scanner, in, "test.Message")

	want := &Oneof{
		Qualified: &Qualified{
			Qualifier: "test.Message.test_oneof",
			Name:      "test_oneof",
		},
		Attributes: []*Attribute{
			{
				Qualified: &Qualified{
					Qualifier: "test.Message",
					Name:      "name",
					Comment:   "Oneof field 1",
				},
				Repeated:    false,
				Optional:    false,
				Map:         false,
				Oneof:       true,
				OneofGroup:  "test_oneof",
				Kind:        []string{"string"},
				Ordinal:     1,
				Annotations: make([]*Annotation, 0),
			},
			{
				Qualified: &Qualified{
					Qualifier: "test.Message",
					Name:      "id",
					Comment:   "Oneof field 2",
				},
				Repeated:    false,
				Optional:    false,
				Map:         false,
				Oneof:       true,
				OneofGroup:  "test_oneof",
				Kind:        []string{"int32"},
				Ordinal:     2,
				Annotations: make([]*Annotation, 0),
			},
		},
	}

	assert.Equal(t, want, got)
}

func TestMessageWithOneofAndFieldsAfter(t *testing.T) {
	testFile := `
  string before_oneof = 1;
  oneof my_oneof {
    string name = 2;
    int32 id = 3;
  }
  string after_oneof = 4;
`
	scanner := NewTestScanner(testFile)
	mv := &MessageVisitor{}
	in := NewLine("message SampleMessage {")

	got := mv.Visit(scanner, in, "test").(*Message)

	assert.Equal(t, "SampleMessage", got.Name)
	assert.Len(t, got.Attributes, 4)

	assert.Equal(t, "before_oneof", got.Attributes[0].Name)
	assert.False(t, got.Attributes[0].Oneof)

	assert.Equal(t, "name", got.Attributes[1].Name)
	assert.True(t, got.Attributes[1].Oneof)
	assert.Equal(t, "my_oneof", got.Attributes[1].OneofGroup)

	assert.Equal(t, "id", got.Attributes[2].Name)
	assert.True(t, got.Attributes[2].Oneof)
	assert.Equal(t, "my_oneof", got.Attributes[2].OneofGroup)

	assert.Equal(t, "after_oneof", got.Attributes[3].Name)
	assert.False(t, got.Attributes[3].Oneof)
}
