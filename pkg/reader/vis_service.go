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
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/api"
	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/logging"
)

type ServiceVisitor struct {
	log      *logging.Logger
	visitors []Visitor
}

func NewServiceVisitor(debug bool) *ServiceVisitor {
	visitors := make([]Visitor, 0)
	visitors = append(visitors, NewRpcVisitor(debug), &CommentVisitor{})
	return &ServiceVisitor{visitors: visitors, log: logging.NewLogger(debug, "service visitor")}
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
		if line.Token == ClosedBrace {
			sv.log.Debug("Finished Reading Service")
			break
		}
	}
	return out
}

var RpcLinePattern = `rpc\s+(.*?)\((.*?)\)\s+returns\s+\((.*?)\)(.*)`

type RpcVisitor struct {
	Log            *logging.Logger
	Visitors       []Visitor
	RpcLineMatcher *regexp.Regexp
}

func NewRpcVisitor(debug bool) *RpcVisitor {
	return &RpcVisitor{
		RpcLineMatcher: regexp.MustCompile(RpcLinePattern),
		Log:            logging.NewLogger(debug, "rpc visitor"),
	}
}

func (rv *RpcVisitor) CanVisit(line *Line) bool {
	rv.Log.Debugf("Checking: %s - Status : %v", line.Syntax, rv.RpcLineMatcher.MatchString(line.Syntax))
	return rv.RpcLineMatcher.MatchString(line.Syntax)
}

func ParseInArgs(values []string, rpc api.RPC) {
	inArgs := strings.Split(values[2], Comma)
	for _, i := range inArgs {
		if strings.HasPrefix(i, "stream") {
			rpc.AddInputParameter(true, strings.TrimSpace(i[strings.Index(i, Space):]))
		} else {
			rpc.AddInputParameter(false, strings.TrimSpace(i))
		}
	}
}

func ParseReturnArgs(values []string, rpc api.RPC) {
	returnArgs := strings.Split(values[3], Comma)
	for _, i := range returnArgs {
		if strings.HasPrefix(i, "stream") {
			rpc.AddReturnParameter(true, strings.TrimSpace(i[strings.Index(i, Space):]))
		} else {
			rpc.AddReturnParameter(false, strings.TrimSpace(i))
		}
	}
}

func (rv *RpcVisitor) Visit(scanner Scanner, in *Line, namespace string) interface{} {
	rv.Log.Debugf("Visiting RPC: %v\n", in.Syntax)

	values := rv.RpcLineMatcher.FindStringSubmatch(in.Syntax)
	out := ProtobufFactory.NewRPC(namespace, values[1], in.Comment)
	ParseInArgs(values, out)
	ParseReturnArgs(values, out)

	rv.Log.Debug("RPC.00 - Scanning RPC")
	for scanner.Scan() {
		line := scanner.ReadLine()
		rv.Log.Debugf("RPC.01: %v with token: %v", line.Syntax, line.Token)

		if strings.HasPrefix(line.Syntax, "option") {
			optionName := line.Syntax[strings.Index(line.Syntax, "(")+1 : strings.Index(line.Syntax, ")")]
			optionBody := ""
			rv.Log.Debug("RPC.02: Reading Body")
			for scanner.Scan() {
				oBody := scanner.ReadLine()
				rv.Log.Debugf("RPC.03: %s with token: %s", oBody.Syntax, oBody.Token)
				optionBody += oBody.Syntax

				if line.Token == Semicolon {
					rv.Log.Debug("Closing Option")
					if len(strings.TrimSpace(optionBody)) > 0 {
						rv.Log.Debugf("Adding Option - %s with body: %s", optionName, optionBody)
						out.AddRPCOption(
							optionName,
							"",
							optionBody)
					}
					break
				}
			}
			rv.Log.Debug("RPC.03: Finished Reading Body")
			if line.Token == ClosedBrace {
				rv.Log.Debugf("Closing RPC")
				break
			}
		}

	}
	return out
}
