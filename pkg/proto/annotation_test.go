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
		// note that even if the source file declares the annotation with white space around `=` some pre-processor upstream of the annotation parser strips it
		{name: "Test without whitespace", args: args{in: "optional uint32 weight = 18 [deprecated=true];"}, want: []*Annotation{{Name: "deprecated", Value: "true"}}},
		// projects that import google/protobuf/timestamp.proto end up parsing the large comment block for annotation and runs into [toISOString()]. There must be another bug upstream, but the Annotation parser shall be protected too.
		{name: "Test google.protobuf.timestamp.proto", args: args{in: "...using the // standard // [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString) // method"}, want: []*Annotation{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ParseAnnotations(tt.args.in), "ParseAnnotations(%v)", tt.args.in)
		})
	}
}
