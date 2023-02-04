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

func TestNamedValue_GetAnchor(t *testing.T) {
	type fields struct {
		Name    string
		Value   string
		Comment Comment
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Get Anchor", fields{
			Name:    "SomeName",
			Value:   "Some Value",
			Comment: "Some Comment",
		}, "some_name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			namedValue := &NamedValue{
				Name:    tt.fields.Name,
				Value:   tt.fields.Value,
				Comment: tt.fields.Comment,
			}
			assert.Equalf(t, tt.want, namedValue.GetAnchor(), "GetAnchor()")
		})
	}
}
