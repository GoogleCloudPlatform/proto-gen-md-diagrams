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

func TestNewEnumValue(t *testing.T) {
	type args struct {
		namespace string
		ordinal   string
		value     string
		comment   Comment
	}
	tests := []struct {
		name string
		args args
		want *EnumValue
	}{
		{name: "Test Enum Value", args: args{
			namespace: "test",
			ordinal:   "1",
			value:     "TEST",
			comment:   "Test",
		}, want: NewEnumValue("test", "1", "TEST", "Test")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewEnumValue(tt.args.namespace, tt.args.ordinal, tt.args.value, tt.args.comment), "NewEnumValue(%v, %v, %v, %v)", tt.args.namespace, tt.args.ordinal, tt.args.value, tt.args.comment)
		})
	}
}
