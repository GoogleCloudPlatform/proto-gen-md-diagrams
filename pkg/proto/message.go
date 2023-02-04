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

// Message represents a message / struct body
type Message struct {
	*Qualified
	Attributes []*Attribute
	Messages   []*Message
	Enums      []*Enum
	Reserved   []*Reserved
}

// NewMessage creates a new message
func NewMessage() *Message {
	return &Message{
		Qualified:  &Qualified{},
		Attributes: make([]*Attribute, 0),
		Messages:   make([]*Message, 0),
		Enums:      make([]*Enum, 0),
		Reserved:   make([]*Reserved, 0),
	}
}

func (m *Message) HasAttributes() bool {
	return len(m.Attributes) > 0
}

func (m *Message) HasMessages() bool {
	return len(m.Messages) > 0
}

func (m *Message) HasEnums() bool {
	return len(m.Enums) > 0
}
