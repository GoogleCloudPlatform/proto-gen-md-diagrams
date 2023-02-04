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

func TestNewService(t *testing.T) {
	type args struct {
		namespace string
		name      string
		comment   Comment
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		{name: "New Service", args: args{
			namespace: "test",
			name:      "TestService",
			comment:   "A Test Service",
		}, want: NewService("test", "TestService", Comment("A Test Service"))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewService(tt.args.namespace, tt.args.name, tt.args.comment), "NewService(%v, %v, %v)", tt.args.namespace, tt.args.name, tt.args.comment)
		})
	}
}

func TestService_AddRpc(t *testing.T) {
	type fields struct {
		Qualified *Qualified
		Methods   []*Rpc
	}
	type args struct {
		rpc []*Rpc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Add RPC", fields: fields{
			Qualified: &Qualified{
				Qualifier: "test",
				Name:      "List",
				Comment:   "List returns a list of physical locations",
			},
			Methods: []*Rpc{},
		}, args: args{rpc: []*Rpc{NewRpc("test", "List", "List returns a list of physical locations")}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Qualified: tt.fields.Qualified,
				Methods:   tt.fields.Methods,
			}
			s.AddRpc(tt.args.rpc...)
			assert.Equal(t, 1, len(s.Methods))
		})
	}
}
