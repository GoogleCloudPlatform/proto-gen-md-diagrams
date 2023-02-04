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

func TestNewAnnotation(t *testing.T) {
	type args struct {
		name  string
		value any
	}
	tests := []struct {
		name string
		args args
		want *Annotation
	}{
		{name: "test 001", args: args{name: "test", value: "test"}, want: &Annotation{"test", "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewAnnotation(tt.args.name, tt.args.value), "NewAnnotation(%v, %v)", tt.args.name, tt.args.value)
		})
	}
}

func TestParseAnnotations(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want []*Annotation
	}{
		{name: "Test 001", args: args{in: "int32 longitude_degrees = 3 [json_name = 'lng_d'];"}, want: []*Annotation{{Name: "json_name", Value: "lng_d"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ParseAnnotations(tt.args.in), "ParseAnnotations(%v)", tt.args.in)
		})
	}
}
