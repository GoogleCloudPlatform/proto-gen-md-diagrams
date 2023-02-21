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

package reader

import (
  "strings"

  "github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/api"
  "github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/logging"
)

type ServiceVisitor struct {
  log      *logging.Logger
  visitors []Visitor
}

func NewServiceVisitor(debug bool) *ServiceVisitor {
  return &ServiceVisitor{
    log:      logging.NewLogger(debug, "service visitor"),
    visitors: []Visitor{&CommentVisitor{}, NewRpcVisitor(debug)}}
}

func (sv *ServiceVisitor) CanVisit(line *Line) bool {
  return strings.HasPrefix(line.Syntax, "service") && line.Token == OpenBrace
}

func (sv *ServiceVisitor) Visit(scanner Scanner, in *Line, namespace string) interface{} {
  sv.log.Debugf("Visiting Service: %v\n", in)

  values := SplitSyntax(in.Syntax)
  out := ProtobufFactory.NewService(namespace, values[1], in.Comment)

  comment := ""
  for scanner.Scan() {
    line := scanner.ReadLine()
    if line.Token == ClosedBrace {
      break
    }
    sv.log.Debugf("Scanning line in service: %s", line.Syntax)
    for _, visitor := range sv.visitors {
      if visitor.CanVisit(line) {
        rt := visitor.Visit(scanner, line, Join(Period, namespace, out.Name()))
        switch t := rt.(type) {
        case api.RPC:
          t.SetComment(Join(Space, comment, t.Comment()))
          out.AddRPC(t)
          comment = ""
        case string:
          comment = Join(Space, comment, t)
        }
      }
    }
  }
  return out
}
