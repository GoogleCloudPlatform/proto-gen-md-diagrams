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

func TestMessage_HasAttributes(t *testing.T) {
	type fields struct {
		Attributes []*Attribute
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "Has Attribute",
			fields: fields{Attributes: []*Attribute{
				NewAttribute("test", "None"),
			}},
			want: true},
		{name: "Doesn't Have Attributes",
			fields: fields{Attributes: make([]*Attribute, 0)},
			want:   false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Attributes: tt.fields.Attributes,
			}
			assert.Equalf(t, tt.want, m.HasAttributes(), "HasAttributes()")
		})
	}
}

func TestMessage_HasEnums(t *testing.T) {
	type fields struct {
		Enums []*Enum
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "Has Enums",
			fields: fields{
				Enums: []*Enum{NewEnum("test", "test", "test")}},
			want: true},
		{name: "Doesn't Have Enums",
			fields: fields{
				Enums: make([]*Enum, 0)},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Enums: tt.fields.Enums,
			}
			assert.Equalf(t, tt.want, m.HasEnums(), "HasEnums()")
		})
	}
}

func TestMessage_HasMessages(t *testing.T) {
	type fields struct {
		Messages []*Message
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "Has Messages", fields: fields{Messages: []*Message{NewMessage()}},
			want: true},
		{name: "Doesn't have Messages", fields: fields{Messages: make([]*Message, 0)},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Messages: tt.fields.Messages,
			}
			assert.Equalf(t, tt.want, m.HasMessages(), "HasMessages()")
		})
	}
}

func TestNewMessage(t *testing.T) {
	tests := []struct {
		name string
		want *Message
	}{
		{name: "New Message", want: &Message{
			Qualified:  &Qualified{},
			Attributes: make([]*Attribute, 0),
			Messages:   make([]*Message, 0),
			Enums:      make([]*Enum, 0),
			Reserved:   make([]*Reserved, 0),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewMessage(), "NewMessage()")
		})
	}
}
