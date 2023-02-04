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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPackage(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *Package
	}{
		{name: "New Package", args: args{path: "test.proto"}, want: &Package{
			Path:     "test.proto",
			Name:     "",
			Comment:  "",
			Options:  make([]*Option, 0),
			Imports:  make([]*Import, 0),
			Messages: make([]*Message, 0),
			Enums:    make([]*Enum, 0),
			Services: make([]*Service, 0),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewPackage(tt.args.path), "NewPackage(%v)", tt.args.path)
		})
	}
}

func TestPackage_Read(t *testing.T) {
	type fields struct {
		Path     string
		Name     string
		Comment  Comment
		Options  []*Option
		Imports  []*Import
		Messages []*Message
		Enums    []*Enum
		Services []*Service
	}
	type args struct {
		debug bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Package{
				Path:     tt.fields.Path,
				Name:     tt.fields.Name,
				Comment:  tt.fields.Comment,
				Options:  tt.fields.Options,
				Imports:  tt.fields.Imports,
				Messages: tt.fields.Messages,
				Enums:    tt.fields.Enums,
				Services: tt.fields.Services,
			}
			tt.wantErr(t, p.Read(tt.args.debug), fmt.Sprintf("Read(%v)", tt.args.debug))
		})
	}
}

func TestPackage_ToMarkdownWithDiagram(t *testing.T) {
	type fields struct {
		Path     string
		Name     string
		Comment  Comment
		Options  []*Option
		Imports  []*Import
		Messages []*Message
		Enums    []*Enum
		Services []*Service
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Package{
				Path:     tt.fields.Path,
				Name:     tt.fields.Name,
				Comment:  tt.fields.Comment,
				Options:  tt.fields.Options,
				Imports:  tt.fields.Imports,
				Messages: tt.fields.Messages,
				Enums:    tt.fields.Enums,
				Services: tt.fields.Services,
			}
			assert.Equalf(t, tt.want, p.ToMarkdownWithDiagram(), "ToMarkdownWithDiagram()")
		})
	}
}
