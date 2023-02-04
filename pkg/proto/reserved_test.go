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

func TestNewReserved(t *testing.T) {
	type args struct {
		start int32
		end   int32
	}
	tests := []struct {
		name string
		args args
		want *Reserved
	}{
		{name: "Reserved", args: args{
			start: 4,
			end:   10,
		}, want: &Reserved{
			Start: 4,
			End:   10,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewReserved(tt.args.start, tt.args.end), "NewReserved(%v, %v)", tt.args.start, tt.args.end)
		})
	}
}
