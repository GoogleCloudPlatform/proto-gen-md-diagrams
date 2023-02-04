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

func TestNewEnum(t *testing.T) {
	type args struct {
		q       string
		name    string
		comment Comment
	}
	tests := []struct {
		name string
		args args
		want *Enum
	}{
		{name: "Test Enum", args: args{q: "test", name: "TEST", comment: "Test"}, want: &Enum{
			Qualified: &Qualified{
				Qualifier: "test",
				Name:      "TEST",
				Comment:   "Test",
			},
			Values: []*EnumValue{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewEnum(tt.args.q, tt.args.name, tt.args.comment), "NewEnum(%v, %v, %v)", tt.args.q, tt.args.name, tt.args.comment)
		})
	}
}
