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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplexExtendedProtobuf(t *testing.T) {
	pkg := NewPackage("data/test/complex/extended_model.proto")
	err := pkg.Read(false)
	assert.Nil(t, err)
	assert.Equal(t, "test.complex", pkg.Name)
	assert.Equal(t, 2, len(pkg.Messages)) // CustomData, ComplexMessage (with nested Result)
	assert.Equal(t, 1, len(pkg.Services))

	config := &WriterConfig{
		visualize:    true,
		pureMarkdown: false,
	}

	markdown := PackageToMarkDown(pkg, config)
	mermaid := PackageToMermaid(pkg)

	// 1. Verify WKT Timestamp produces no relationship line
	assert.NotContains(t, mermaid, "ComplexMessage --> `google.protobuf.Timestamp`", "WKT Timestamp should not produce a relationship arrow")
	assert.NotContains(t, mermaid, "ComplexMessage --> `Timestamp`", "WKT Timestamp should not produce a relationship arrow")

	// 2. Verify oneof payload metadata
	assert.Contains(t, mermaid, "[oneof: payload]", "Mermaid diagram should include oneof group name")
	assert.Contains(t, markdown, "Oneof (payload)", "Markdown table should include Oneof (payload)")

	// 3. Verify deprecated field annotations
	assert.Contains(t, mermaid, "[deprecated]", "Mermaid diagram should include deprecated annotation tag")
	assert.Contains(t, markdown, "Deprecated", "Markdown table should include Deprecated label")

	// 4. Verify Proto2 group processing
	// ComplexMessage should contain a nested message "Result" and attribute "result"
	complexMsg := pkg.Messages[1]
	assert.Equal(t, "ComplexMessage", complexMsg.Name)
	assert.Equal(t, 1, len(complexMsg.Messages))
	assert.Equal(t, "Result", complexMsg.Messages[0].Name)

	hasResultAttr := false
	for _, attr := range complexMsg.Attributes {
		if attr.Name == "result" && strings.Contains(attr.Kind[0], "Result") {
			hasResultAttr = true
			break
		}
	}
	assert.True(t, hasResultAttr, "ComplexMessage should contain a field attribute for the Proto2 group 'result'")

	// 5. Verify Service streaming relationships
	assert.Contains(t, mermaid, "ComplexService --o `CustomData` : client stream", "Mermaid should render client stream relationship")
	assert.Contains(t, mermaid, "ComplexService --o `ComplexMessage` : server stream", "Mermaid should render server stream relationship")
}
