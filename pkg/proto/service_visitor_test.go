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

func TestNewServiceVisitor(t *testing.T) {
	tests := []struct {
		name string
		want *ServiceVisitor
	}{
		{name: "New Service Visitor", want: NewServiceVisitor()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewServiceVisitor(), "NewServiceVisitor()")
		})
	}
}

func TestServiceVisitor_CanVisit(t *testing.T) {
	type fields struct {
		Visitors []Visitor
	}
	type args struct {
		line *Line
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "Can Visit", fields: fields{Visitors: []Visitor{}}, args: args{&Line{Syntax: "service LocationService", Token: OpenBrace}}, want: true},
		{name: "Can't Visit", fields: fields{Visitors: []Visitor{}}, args: args{&Line{Syntax: "A Comment", Token: InlineCommentPrefix}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv := &ServiceVisitor{
				Visitors: tt.fields.Visitors,
			}
			assert.Equalf(t, tt.want, sv.CanVisit(tt.args.line), "CanVisit(%v)", tt.args.line)
		})
	}
}

func TestServiceVisitor_Visit(t *testing.T) {
	type fields struct {
		Visitors []Visitor
	}
	type args struct {
		scanner   Scanner
		in        *Line
		namespace string
	}

	scanner := NewTestScanner(`
  // List returns a list of physical locations
  rpc List(google.protobuf.Empty) returns (stream test.location.PhysicalLocation) {
      // Creates the get location
      option (google.api.http) = {
        get: "/locations"
      };
  }
`)

	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{name: "Service Visit", fields: fields{Visitors: []Visitor{NewRpcVisitor()}}, args: args{
			scanner:   scanner,
			in:        &Line{Syntax: "service LocationService", Token: OpenBrace},
			namespace: "test.service",
		}, want: &Service{
			Qualified: &Qualified{
				Qualifier: "test.service",
				Name:      "LocationService",
				Comment:   "",
			},
			Methods: []*Rpc{
				{
					Qualified: &Qualified{
						Qualifier: "test.service.LocationService",
						Name:      "List",
						Comment:   "",
					},
					InputParameters:  []*Parameter{NewParameter(false, "google.protobuf.Empty")},
					ReturnParameters: []*Parameter{NewParameter(true, "test.location.PhysicalLocation")},
					Options:          []*RpcOption{NewRpcOption("test.service.LocationService.List", "google.api.http", "", "}")},
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv := &ServiceVisitor{
				Visitors: tt.fields.Visitors,
			}
			assert.Equalf(t, tt.want, sv.Visit(tt.args.scanner, tt.args.in, tt.args.namespace), "Visit(%v, %v, %v)", tt.args.scanner, tt.args.in, tt.args.namespace)
		})
	}
}
