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

func TestEnumToMarkdown(t *testing.T) {
	type args struct {
		enum *Enum
		wc   *WriterConfig
	}
	tests := []struct {
		name        string
		args        args
		wantBody    string
		wantDiagram string
	}{
		{name: "Enum Markdown", args: args{
			enum: &Enum{
				Qualified: &Qualified{
					Qualifier: "test.TestEnum",
					Name:      "TestEnum",
					Comment:   "Keen Enum",
				},
				Values: []*EnumValue{
					NewEnumValue("test.TestEnum", "0", "T_01", ""),
					NewEnumValue("test.TestEnum", "1", "T_02", ""),
				},
			},
			wc: &WriterConfig{
				visualize: false,
			},
		}, wantBody: `## Enum: TestEnum
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: test.TestEnum</div>

<div class="comment"><span>Keen Enum</span><br/></div>

| Name | Ordinal | Description |
|------|---------|-------------|
| T_01 | 0       |             |
| T_02 | 1       |             |


`, wantDiagram: ``},
		{name: "Enum Markdown", args: args{
			enum: &Enum{
				Qualified: &Qualified{
					Qualifier: "test.TestEnum",
					Name:      "TestEnum",
					Comment:   "Keen Enum",
				},
				Values: []*EnumValue{
					NewEnumValue("test.TestEnum", "0", "T_01", ""),
					NewEnumValue("test.TestEnum", "1", "T_02", ""),
				},
			},
			wc: &WriterConfig{
				visualize:    false,
				pureMarkdown: true,
			},
		}, wantBody: `## Enum: TestEnum
* **FQN**: test.TestEnum

Keen Enum 

| Name | Ordinal | Description |
|------|---------|-------------|
| T_01 | 0       |             |
| T_02 | 1       |             |


`, wantDiagram: ``},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBody, gotDiagram := EnumToMarkdown(tt.args.enum, tt.args.wc)
			assert.Equalf(t, tt.wantBody, gotBody, "EnumToMarkdown(%v, %v)", tt.args.enum, tt.args.wc)
			assert.Equalf(t, tt.wantDiagram, gotDiagram, "EnumToMarkdown(%v, %v)", tt.args.enum, tt.args.wc)
		})
	}
}

func TestFormatServiceParameter(t *testing.T) {
	type args struct {
		parameters []*Parameter
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Service Parameter",
			args: args{parameters: []*Parameter{NewParameter(false, "test.location.PhysicalLocation")}},
			want: "PhysicalLocation"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FormatServiceParameter(tt.args.parameters), "FormatServiceParameter(%v)", tt.args.parameters)
		})
	}
}

func TestHandleEnums(t *testing.T) {
	type args struct {
		enums []*Enum
		wc    *WriterConfig
	}
	tests := []struct {
		name     string
		args     args
		wantBody string
	}{
		{name: "Enums", args: args{
			enums: []*Enum{{
				Qualified: &Qualified{
					Qualifier: "test.Service",
					Name:      "TestEnum",
					Comment:   "",
				},
				Values: []*EnumValue{NewEnumValue("test.Service.TestEnum", "0", "T1", "")},
			}},
			wc: &WriterConfig{
				visualize: false,
			},
		}, wantBody: `## Enum: TestEnum
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: test.Service</div>

<div class="comment"><span></span><br/></div>

| Name | Ordinal | Description |
|------|---------|-------------|
| T1   | 0       |             |


`},
		{name: "Enums", args: args{
			enums: []*Enum{{
				Qualified: &Qualified{
					Qualifier: "test.Service",
					Name:      "TestEnum",
					Comment:   "",
				},
				Values: []*EnumValue{NewEnumValue("test.Service.TestEnum", "0", "T1", "")},
			}},
			wc: &WriterConfig{
				visualize:    false,
				pureMarkdown: true,
			},
		}, wantBody: `## Enum: TestEnum
* **FQN**: test.Service

 

| Name | Ordinal | Description |
|------|---------|-------------|
| T1   | 0       |             |


`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantBody, HandleEnums(tt.args.enums, tt.args.wc), "HandleEnums(%v, %v)", tt.args.enums, tt.args.wc)
		})
	}
}

func TestHandleMessages(t *testing.T) {
	type args struct {
		messages []*Message
		wc       *WriterConfig
	}
	tests := []struct {
		name     string
		args     args
		wantBody string
	}{
		{name: "Handle Message", args: args{
			messages: []*Message{&Message{
				Qualified: &Qualified{
					Qualifier: "test.Service.Message",
					Name:      "Message",
					Comment:   "",
				},
				Attributes: []*Attribute{{
					Qualified: &Qualified{
						Qualifier: "test.Service.Message.Name",
						Name:      "Name",
						Comment:   "",
					},
					Repeated:    false,
					Map:         false,
					Kind:        []string{"string"},
					Ordinal:     1,
					Annotations: []*Annotation{},
				}},
				Messages: []*Message{},
				Enums:    []*Enum{},
				Reserved: []*Reserved{},
			}},
			wc: &WriterConfig{
				visualize: false,
			},
		}, wantBody: `## Message: Message
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: test.Service.Message</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type   | Label | Description |
|-------|---------|--------|-------|-------------|
| Name  | 1       | string |       |             |


`},
		{name: "Handle Message", args: args{
			messages: []*Message{&Message{
				Qualified: &Qualified{
					Qualifier: "test.Service.Message",
					Name:      "Message",
					Comment:   "",
				},
				Attributes: []*Attribute{{
					Qualified: &Qualified{
						Qualifier: "test.Service.Message.Name",
						Name:      "Name",
						Comment:   "",
					},
					Repeated:    false,
					Map:         false,
					Kind:        []string{"string"},
					Ordinal:     1,
					Annotations: []*Annotation{},
				}},
				Messages: []*Message{},
				Enums:    []*Enum{},
				Reserved: []*Reserved{},
			}},
			wc: &WriterConfig{
				visualize:    false,
				pureMarkdown: true,
			},
		}, wantBody: `## Message: Message
* **FQN**: test.Service.Message

 

| Field | Ordinal | Type   | Label | Description |
|-------|---------|--------|-------|-------------|
| Name  | 1       | string |       |             |


`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantBody, HandleMessages(tt.args.messages, tt.args.wc), "HandleMessages(%v, %v)", tt.args.messages, tt.args.wc)
		})
	}
}

func TestMessageToMarkdown(t *testing.T) {
	type args struct {
		message *Message
		wc      *WriterConfig
	}
	tests := []struct {
		name        string
		args        args
		wantBody    string
		wantDiagram string
	}{
		{name: "Handle Message", args: args{
			message: &Message{
				Qualified: &Qualified{
					Qualifier: "test.Service.Message",
					Name:      "Message",
					Comment:   "",
				},
				Attributes: []*Attribute{{
					Qualified: &Qualified{
						Qualifier: "test.Service.Message.Name",
						Name:      "Name",
						Comment:   "",
					},
					Repeated:    false,
					Map:         false,
					Kind:        []string{"string"},
					Ordinal:     1,
					Annotations: []*Annotation{},
				}},
				Messages: []*Message{},
				Enums:    []*Enum{},
				Reserved: []*Reserved{},
			},
			wc: &WriterConfig{
				visualize: true,
			},
		}, wantBody: `## Message: Message
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: test.Service.Message</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type   | Label | Description |
|-------|---------|--------|-------|-------------|
| Name  | 1       | string |       |             |


`, wantDiagram: "\n### Message Diagram\n\n```mermaid\nclassDiagram\ndirection LR\n\n%% \n\nclass Message {\n  + string Name\n}\n\n```"},
		{name: "Handle Message", args: args{
			message: &Message{
				Qualified: &Qualified{
					Qualifier: "test.Service.Message",
					Name:      "Message",
					Comment:   "",
				},
				Attributes: []*Attribute{{
					Qualified: &Qualified{
						Qualifier: "test.Service.Message.Name",
						Name:      "Name",
						Comment:   "",
					},
					Repeated:    false,
					Map:         false,
					Kind:        []string{"string"},
					Ordinal:     1,
					Annotations: []*Annotation{},
				}},
				Messages: []*Message{},
				Enums:    []*Enum{},
				Reserved: []*Reserved{},
			},
			wc: &WriterConfig{
				visualize:    true,
				pureMarkdown: true,
			},
		}, wantBody: `## Message: Message
* **FQN**: test.Service.Message

 

| Field | Ordinal | Type   | Label | Description |
|-------|---------|--------|-------|-------------|
| Name  | 1       | string |       |             |


`, wantDiagram: "\n### Message Diagram\n\n```mermaid\nclassDiagram\ndirection LR\n\n%% \n\nclass Message {\n  + string Name\n}\n\n```"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBody, gotDiagram := MessageToMarkdown(tt.args.message, tt.args.wc)
			assert.Equalf(t, tt.wantBody, gotBody, "MessageToMarkdown(%v, %v)", tt.args.message, tt.args.wc)
			assert.Equalf(t, tt.wantDiagram, gotDiagram, "MessageToMarkdown(%v, %v)", tt.args.message, tt.args.wc)
		})
	}
}

func TestPackageFormatImports(t *testing.T) {
	type args struct {
		p *Package
	}
	tests := []struct {
		name     string
		args     args
		wantBody string
	}{
		{name: "Package Format", args: args{p: &Package{
			Path:    "/some/file/path/example.proto",
			Name:    "test.package",
			Comment: "A test Package",
			Imports: []*Import{{
				Path:    "test/location/model.proto",
				Comment: "None",
			}},
		}}, wantBody: `## Imports

| Import                    | Description |
|---------------------------|-------------|
| test/location/model.proto | None        |

`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantBody, PackageFormatImports(tt.args.p), "PackageFormatImports(%v)", tt.args.p)
		})
	}
}

func TestPackageFormatOptions(t *testing.T) {
	type args struct {
		p *Package
	}
	tests := []struct {
		name     string
		args     args
		wantBody string
	}{
		{
			name: "Format Options",
			args: args{p: &Package{
				Path:    "test/location/model.proto",
				Name:    "test.package",
				Comment: "None",
				Options: []*Option{&Option{NamedValue: &NamedValue{
					Name:    "go_package",
					Value:   "gcp/proto/test/location",
					Comment: "",
				}}},
			}},
			wantBody: `## Options

| Name       | Value                   | Description |
|------------|-------------------------|-------------|
| go_package | gcp/proto/test/location |             |

`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantBody, PackageFormatOptions(tt.args.p), "PackageFormatOptions(%v)", tt.args.p)
		})
	}
}

func TestPackageToMarkDown(t *testing.T) {
	type args struct {
		p  *Package
		wc *WriterConfig
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Package", args: args{
			p: &Package{
				Path:     "test/location/model.proto",
				Name:     "test.package",
				Comment:  "",
				Options:  []*Option{},
				Imports:  []*Import{},
				Messages: []*Message{},
				Enums:    []*Enum{},
				Services: []*Service{},
			},
			wc: &WriterConfig{},
		}, want: `# Package: test.package

<div class="comment"><span></span><br/></div>

## Imports

| Import | Description |
|--------|-------------|



## Options

| Name | Value | Description |
|------|-------|-------------|





<!-- Created by: Proto Diagram Tool -->
<!-- https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams -->
`},
		{name: "Package", args: args{
			p: &Package{
				Path:     "test/location/model.proto",
				Name:     "test.package",
				Comment:  "",
				Options:  []*Option{},
				Imports:  []*Import{},
				Messages: []*Message{},
				Enums:    []*Enum{},
				Services: []*Service{},
			},
			wc: &WriterConfig{
				pureMarkdown: true,
			},
		}, want: `# Package: test.package

 

## Imports

| Import | Description |
|--------|-------------|



## Options

| Name | Value | Description |
|------|-------|-------------|





<!-- Created by: Proto Diagram Tool -->
<!-- https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams -->
`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PackageToMarkDown(tt.args.p, tt.args.wc), "PackageToMarkDown(%v, %v)", tt.args.p, tt.args.wc)
		})
	}
}

func TestServiceToMarkdown(t *testing.T) {
	type args struct {
		s  *Service
		wc *WriterConfig
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Service", args: args{s: &Service{
			Qualified: &Qualified{
				Qualifier: "test.Service",
				Name:      "Service",
				Comment:   "",
			},
			Methods: []*Rpc{},
		}, wc: &WriterConfig{}}, want: `## Service: Service
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: test.Service</div>

<div class="comment"><span></span><br/></div>

| Method | Parameter (In) | Parameter (Out) | Description |
|--------|----------------|-----------------|-------------|


`},
		{name: "Service", args: args{s: &Service{
			Qualified: &Qualified{
				Qualifier: "test.Service",
				Name:      "Service",
				Comment:   "",
			},
			Methods: []*Rpc{},
		}, wc: &WriterConfig{
			pureMarkdown: true,
		}}, want: `## Service: Service
* **FQN**: test.Service

 

| Method | Parameter (In) | Parameter (Out) | Description |
|--------|----------------|-----------------|-------------|


`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ServiceToMarkdown(tt.args.s, tt.args.wc), "ServiceToMarkdown(%v, %v)", tt.args.s, tt.args.wc)
		})
	}
}
