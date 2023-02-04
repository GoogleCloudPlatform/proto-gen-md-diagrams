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

func TestAttribute_IsValid(t *testing.T) {
	type fields struct {
		Qualified   *Qualified
		Repeated    bool
		Map         bool
		Kind        []string
		Ordinal     int
		Annotations []*Annotation
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "Test 001", fields: fields{Qualified: &Qualified{
			Qualifier: "test.qualifier",
			Name:      "Test",
			Comment:   "This is a test",
		}, Repeated: false, Map: false, Kind: []string{"string"}, Ordinal: 1}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Attribute{
				Qualified:   tt.fields.Qualified,
				Repeated:    tt.fields.Repeated,
				Map:         tt.fields.Map,
				Kind:        tt.fields.Kind,
				Ordinal:     tt.fields.Ordinal,
				Annotations: tt.fields.Annotations,
			}
			assert.Equalf(t, tt.want, a.IsValid(), "IsValid()")
		})
	}
}

func TestAttribute_ToMermaid(t *testing.T) {
	type fields struct {
		Qualified   *Qualified
		Repeated    bool
		Map         bool
		Kind        []string
		Ordinal     int
		Annotations []*Annotation
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Test 001", fields: fields{Qualified: &Qualified{
			Qualifier: "test.qualifier",
			Name:      "Test",
			Comment:   "This is a test",
		}, Repeated: false, Map: false, Kind: []string{"string"}, Ordinal: 1}, want: "+ string Test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Attribute{
				Qualified:   tt.fields.Qualified,
				Repeated:    tt.fields.Repeated,
				Map:         tt.fields.Map,
				Kind:        tt.fields.Kind,
				Ordinal:     tt.fields.Ordinal,
				Annotations: tt.fields.Annotations,
			}
			assert.Equalf(t, tt.want, a.ToMermaid(), "ToMermaid()")
		})
	}
}

func TestNewAttribute(t *testing.T) {
	type args struct {
		namespace string
		comment   Comment
	}
	tests := []struct {
		name string
		args args
		want *Attribute
	}{
		{name: "Test 001", args: args{namespace: "test.namespace", comment: Comment("testing")}, want: &Attribute{
			Qualified: &Qualified{
				Qualifier: "test.namespace",
				Name:      "",
				Comment:   "testing",
			}, Annotations: make([]*Annotation, 0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewAttribute(tt.args.namespace, tt.args.comment), "NewAttribute(%v, %v)", tt.args.namespace, tt.args.comment)
		})
	}
}
