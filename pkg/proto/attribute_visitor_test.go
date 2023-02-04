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

func TestNewAttributeVisitor(t *testing.T) {
	tests := []struct {
		name string
		want *AttributeVisitor
	}{
		{name: "New Visitor", want: &AttributeVisitor{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewAttributeVisitor(), "NewAttributeVisitor()")
		})
	}
}

func Test_attributeVisitor_CanVisit(t *testing.T) {
	type args struct {
		in *Line
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Can Visit", args: args{in: &Line{
			Syntax:  "string line1 = 1",
			Token:   ";",
			Comment: "Test",
		}}, want: true},
		{name: "Can Not Visit Comment", args: args{in: &Line{
			Syntax:  "",
			Token:   "//",
			Comment: "Test",
		}}, want: false},
		{name: "Can Not Visit Message", args: args{in: &Line{
			Syntax:  "message Address",
			Token:   "{",
			Comment: "",
		}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			av := AttributeVisitor{}
			assert.Equalf(t, tt.want, av.CanVisit(tt.args.in), "CanVisit(%v)", tt.args.in)
		})
	}
}

func Test_attributeVisitor_Visit(t *testing.T) {
	type args struct {
		in0       Scanner
		in        *Line
		namespace string
	}

	testScanner := NewTestScanner(``)

	tests := []struct {
		name string
		args args
		want *Attribute
	}{
		{
			name: "Test Visit", args: args{in0: testScanner,
			in:        &Line{Syntax: "string line1 = 1", Token: ";", Comment: "Test"},
			namespace: "test",
		},
			want: &Attribute{
				Qualified: &Qualified{
					Qualifier: "test",
					Name:      "line1",
					Comment:   "Test",
				},
				Repeated:    false,
				Map:         false,
				Kind:        []string{"string"},
				Ordinal:     1,
				Annotations: make([]*Annotation, 0),
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			av := AttributeVisitor{}
			output := av.Visit(tt.args.in0, tt.args.in, tt.args.namespace)
			switch o := output.(type) {
			case *Attribute:
				assert.Equal(t, tt.want.Qualifier, o.Qualifier)
				assert.Equal(t, tt.want.Name, o.Name)
				assert.Equal(t, tt.want.Comment, o.Comment)
			default:
				assert.Fail(t, "Failed to convert type")
			}

		})
	}
}

func Test_handleDefaultAttribute(t *testing.T) {
	type args struct {
		out   *Attribute
		split []string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Handle Default Attribute",
			args: args{out: NewAttribute("test", "Test"), split: []string{"string", "line1", "=", "1"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandleDefaultAttribute(tt.args.out, tt.args.split)
			assert.Equalf(t, "line1", tt.args.out.Name, "Name %s", tt.args.out.Name)
			assert.Equalf(t, 1, tt.args.out.Ordinal, "Name %d", tt.args.out.Ordinal)
		})
	}
}

func Test_HandleMap(t *testing.T) {
	type args struct {
		out   *Attribute
		split []string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Test Map", args: args{
			out:   NewAttribute("test", "Test"),
			split: []string{"map<string,", "string>", "meta", "=", "11"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandleMap(tt.args.out, tt.args.split)
			assert.True(t, tt.args.out.Map)
		})
	}
}

func Test_HandleRepeated(t *testing.T) {
	type args struct {
		out   *Attribute
		split []string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Test Map", args: args{
			out:   NewAttribute("test", "Test"),
			split: []string{"repeated", "string", "name", "=", "1"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandleRepeated(tt.args.out, tt.args.split)
			assert.True(t, tt.args.out.Repeated)
		})
	}
}
