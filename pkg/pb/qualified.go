/*
 * Copyright 2023 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pb

import "github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/api"

type qualified struct {
	qualifier string
	name      string
	comment   string
}

func newQualified(qualifier string, name string, comment string) api.Qualified {
	return &qualified{qualifier: qualifier, name: name, comment: comment}
}

func (q *qualified) Qualifier() string {
	return q.qualifier
}

func (q *qualified) Name() string {
	return q.name
}

func (q *qualified) Comment() string {
	return q.comment
}

func (q *qualified) SetComment(in string) {
	q.comment = in
}
