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
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRpcVisitor(t *testing.T) {
	tests := []struct {
		name string
		want *RpcVisitor
	}{
		{name: "New RPC Visitor", want: NewRpcVisitor()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewRpcVisitor(), "NewRpcVisitor()")
		})
	}
}

func TestRpcVisitor_CanVisit(t *testing.T) {
	type fields struct {
		Visitors       []Visitor
		RpcLineMarcher *regexp.Regexp
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
		{name: "Can Visit", fields: fields{
			Visitors:       make([]Visitor, 0),
			RpcLineMarcher: regexp.MustCompile(RpcLinePattern),
		}, args: args{
			line: &Line{
				Syntax:  "rpc List(google.protobuf.Empty) returns (stream test.location.PhysicalLocation)",
				Token:   OpenBrace,
				Comment: "List returns a list of physical locations",
			},
		}, want: true},
		{name: "Can't Visit", fields: fields{
			Visitors:       make([]Visitor, 0),
			RpcLineMarcher: regexp.MustCompile(RpcLinePattern),
		}, args: args{
			line: &Line{Syntax: "Comment", Token: InlineCommentPrefix},
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rv := &RpcVisitor{
				Visitors:       tt.fields.Visitors,
				RpcLineMatcher: tt.fields.RpcLineMarcher,
			}
			assert.Equalf(t, tt.want, rv.CanVisit(tt.args.line), "CanVisit(%v)", tt.args.line)
		})
	}
}

func TestRpcVisitor_Visit(t *testing.T) {
	type fields struct {
		Visitors       []Visitor
		RpcLineMarcher *regexp.Regexp
	}
	type args struct {
		in        *Line
		namespace string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name: "Visit",
			fields: fields{
				Visitors:       []Visitor{},
				RpcLineMarcher: regexp.MustCompile(RpcLinePattern),
			},
			args: args{
				in: &Line{
					Syntax:  "rpc List(google.protobuf.Empty) returns (stream test.location.PhysicalLocation)",
					Token:   OpenBrace,
					Comment: "List returns a list of physical locations",
				},
				namespace: "test.LocationService",
			},
			want: &Rpc{
				Qualified: &Qualified{
					Qualifier: "test.LocationService",
					Name:      "List",
					Comment:   "List returns a list of physical locations",
				},
				InputParameters:  []*Parameter{NewParameter(false, "google.protobuf.Empty")},
				ReturnParameters: []*Parameter{NewParameter(true, "test.location.PhysicalLocation")},
				Options:          []*RpcOption{NewRpcOption("test.LocationService.List", "google.api.http", "", "}")},
			},
		},
		{
			name: "Visit",
			fields: fields{
				Visitors:       []Visitor{},
				RpcLineMarcher: regexp.MustCompile(RpcLinePattern),
			},
			args: args{
				in: &Line{
					Syntax:  "rpc List(google.protobuf.Empty) returns (test.location.PhysicalLocation)",
					Token:   OpenBrace,
					Comment: "List returns a list of physical locations",
				},
				namespace: "test.LocationService",
			},
			want: &Rpc{
				Qualified: &Qualified{
					Qualifier: "test.LocationService",
					Name:      "List",
					Comment:   "List returns a list of physical locations",
				},
				InputParameters:  []*Parameter{NewParameter(false, "google.protobuf.Empty")},
				ReturnParameters: []*Parameter{NewParameter(false, "test.location.PhysicalLocation")},
				Options:          []*RpcOption{NewRpcOption("test.LocationService.List", "google.api.http", "", "}")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rv := &RpcVisitor{
				Visitors:       tt.fields.Visitors,
				RpcLineMatcher: tt.fields.RpcLineMarcher,
			}
			testScanner := NewTestScanner(`
					// Creates the get location
					option (google.api.http) = {
						get: "/locations"
					};
			}
`)
			assert.Equalf(t, tt.want, rv.Visit(testScanner, tt.args.in, tt.args.namespace), "Visit(%v, %v, %v)", testScanner, tt.args.in, tt.args.namespace)
		})
	}
}
