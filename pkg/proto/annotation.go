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

import "strings"

// Annotation is an inline structure applicable only to attributes
type Annotation struct {
	Name  string
	Value any
}

// NewAnnotation is the Annotation Constructor
func NewAnnotation(name string, value any) *Annotation {
	return &Annotation{Name: name, Value: value}
}

// ParseAnnotations is used for reading the annotation line and marshalling it into
// the annotation structure.
func ParseAnnotations(in string) []*Annotation {
	Log.Debug("Processing Annotation")
	out := make([]*Annotation, 0)
	if strings.Contains(in, OpenBracket) && strings.Contains(in, ClosedBracket) {
		annotationString := in[strings.Index(in, OpenBracket)+1 : strings.Index(in, ClosedBracket)]
		split := strings.Split(strings.ReplaceAll(annotationString, SingleQuote, Empty), Space)
		out = append(out, NewAnnotation(split[0], split[2]))
	}
	return out
}
