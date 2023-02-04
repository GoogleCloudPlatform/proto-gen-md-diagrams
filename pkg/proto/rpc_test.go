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

func TestNewParameter(t *testing.T) {
	type args struct {
		stream bool
		t      string
	}
	tests := []struct {
		name string
		args args
		want *Parameter
	}{
		{name: "New Parameter", args: args{
			stream: false,
			t:      "test",
		}, want: NewParameter(false, "test")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewParameter(tt.args.stream, tt.args.t), "NewParameter(%v, %v)", tt.args.stream, tt.args.t)
		})
	}
}

func TestNewRpc(t *testing.T) {
	type args struct {
		namespace string
		name      string
		comment   Comment
	}
	tests := []struct {
		name string
		args args
		want *Rpc
	}{
		{name: "New RPC", args: args{
			namespace: "test.Service",
			name:      "List",
			comment:   "Test Comment",
		}, want: NewRpc("test.Service", "List", Comment("Test Comment"))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewRpc(tt.args.namespace, tt.args.name, tt.args.comment), "NewRpc(%v, %v, %v)", tt.args.namespace, tt.args.name, tt.args.comment)
		})
	}
}

func TestNewRpcOption(t *testing.T) {
	type args struct {
		namespace string
		name      string
		comment   Comment
		body      string
	}
	tests := []struct {
		name string
		args args
		want *RpcOption
	}{
		{name: "New RPC Option", args: args{
			namespace: "test.Service",
			name:      "TestOption",
			comment:   "Test Comment",
			body:      "test",
		}, want: NewRpcOption("test.Service", "TestOption", "Test Comment", "test")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewRpcOption(tt.args.namespace, tt.args.name, tt.args.comment, tt.args.body), "NewRpcOption(%v, %v, %v, %v)", tt.args.namespace, tt.args.name, tt.args.comment, tt.args.body)
		})
	}
}

func TestRpc_AddInputParameter(t *testing.T) {
	type fields struct {
		Qualified        *Qualified
		InputParameters  []*Parameter
		ReturnParameters []*Parameter
		Options          []*RpcOption
	}
	type args struct {
		params []*Parameter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Add Input Parameter", fields: fields{InputParameters: make([]*Parameter, 0)}, args: args{params: []*Parameter{NewParameter(false, "test")}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpc := &Rpc{
				Qualified:        tt.fields.Qualified,
				InputParameters:  tt.fields.InputParameters,
				ReturnParameters: tt.fields.ReturnParameters,
				Options:          tt.fields.Options,
			}
			rpc.AddInputParameter(tt.args.params...)
			assert.Equal(t, 1, len(rpc.InputParameters))
		})
	}
}

func TestRpc_AddReturnParameter(t *testing.T) {
	type fields struct {
		Qualified        *Qualified
		InputParameters  []*Parameter
		ReturnParameters []*Parameter
		Options          []*RpcOption
	}
	type args struct {
		params []*Parameter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Add Input Parameter", fields: fields{ReturnParameters: make([]*Parameter, 0)}, args: args{params: []*Parameter{NewParameter(false, "test")}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpc := &Rpc{
				Qualified:        tt.fields.Qualified,
				InputParameters:  tt.fields.InputParameters,
				ReturnParameters: tt.fields.ReturnParameters,
				Options:          tt.fields.Options,
			}
			rpc.AddReturnParameter(tt.args.params...)
			assert.Equal(t, 1, len(rpc.ReturnParameters))
		})
	}
}

func TestRpc_AddRpcOption(t *testing.T) {
	type fields struct {
		Qualified        *Qualified
		InputParameters  []*Parameter
		ReturnParameters []*Parameter
		Options          []*RpcOption
	}
	type args struct {
		options []*RpcOption
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Add RPC Option", fields: fields{Options: make([]*RpcOption, 0)}, args: args{options: []*RpcOption{NewRpcOption("test", "test", Comment("test"), "test")}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpc := &Rpc{
				Qualified:        tt.fields.Qualified,
				InputParameters:  tt.fields.InputParameters,
				ReturnParameters: tt.fields.ReturnParameters,
				Options:          tt.fields.Options,
			}
			rpc.AddRpcOption(tt.args.options...)
			assert.Equal(t, 1, len(rpc.Options))
		})
	}
}
