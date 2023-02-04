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
)

// NewAttributeVisitor - Constructor for the AttributeVisitor
func NewAttributeVisitor() *AttributeVisitor {
	Log.Debug("Initializing attributeVisitor")
	return &AttributeVisitor{}
}

// Visitor implementation for attributes
type AttributeVisitor struct {
}

// CanVisit - Determines if the line is an attribute, it doesn't end in a brace,
// it's a map, repeated, or can effectively be split
func (av *AttributeVisitor) CanVisit(in *Line) bool {
	return (!strings.HasSuffix(in.Syntax, OpenBrace) || !strings.HasSuffix(in.Syntax, CloseBrace)) &&
			strings.HasPrefix(in.Syntax, "repeated") ||
			strings.HasPrefix(in.Syntax, "map") || len(in.SplitSyntax()) >= 4
}

// HandleRepeated marshals the attribute into a repeated representation, e.g. List.
func HandleRepeated(out *Attribute, split []string) {
	Log.Debugf("\t processing repeated attribute %s", split[2])
	// 0 - 4 repeated, type, name, equals, ordinal
	out.Repeated = true
	out.Kind = append(out.Kind, split[1])
	out.Name = split[2]
	out.Ordinal = ParseOrdinal(split[4])
}

// handleMap marshals the attribute into a Map type by using multiple types for key and value.
func HandleMap(out *Attribute, split []string) {
	Log.Debugf("\t processing map attribute %s", split[2])
	// map1, map2, name, equals, ordinal
	mapValue := Join(Space, split[0], split[1])
	innerTypes := mapValue[strings.Index(mapValue, OpenMap)+len(OpenMap) : strings.Index(mapValue, CloseMap)]
	splitTypes := strings.Split(innerTypes, Comma)
	out.Name = split[2]
	out.Map = true
	out.Kind = append(out.Kind, splitTypes...)
	out.Ordinal = ParseOrdinal(split[4])
}

// HandleDefaultAttribute marshals a standard attribute type.
func HandleDefaultAttribute(out *Attribute, split []string) {
	if len(split) >= 3 {
		Log.Debugf("\t processing standard attribute %s", split[1])
		out.Name = split[1]
		out.Kind = append(out.Kind, split[0])
		out.Ordinal = ParseOrdinal(split[3])
	}
}

// Visit is used for marshalling an attribute into a struct.
func (av *AttributeVisitor) Visit(_ Scanner, in *Line, namespace string) interface{} {
	Log.Debug("Visiting Attribute")
	out := NewAttribute(namespace, in.Comment)
	out.Annotations = ParseAnnotations(in.Syntax)
	split := in.SplitSyntax()

	if strings.HasPrefix(in.Syntax, PrefixReserved) {
		Log.Debug("\t processing reserved attribute")
		out.Comment += Space + in.Comment
	} else if strings.HasPrefix(in.Syntax, PrefixRepeated) {
		HandleRepeated(out, split)
	} else if strings.HasPrefix(in.Syntax, PrefixMap) {
		HandleMap(out, split)
	} else {
		HandleDefaultAttribute(out, split)
	}
	return out
}
