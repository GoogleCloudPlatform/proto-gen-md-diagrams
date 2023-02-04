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
	"regexp"
	"strings"
)

var RpcLinePattern = `rpc\s+(.*?)\((.*?)\)\s+returns\s+\((.*?)\)(.*)`

type RpcVisitor struct {
	Visitors       []Visitor
	RpcLineMatcher *regexp.Regexp
}

func NewRpcVisitor() *RpcVisitor {
	return &RpcVisitor{RpcLineMatcher: regexp.MustCompile(RpcLinePattern)}
}

func (rv *RpcVisitor) CanVisit(line *Line) bool {
	return rv.RpcLineMatcher.MatchString(line.Syntax)
}

func ParseInArgs(values []string, rpc *Rpc) {
	inArgs := strings.Split(values[2], Comma)
	for _, i := range inArgs {
		if strings.HasPrefix(i, "stream") {
			rpc.AddInputParameter(NewParameter(true, strings.TrimSpace(i[strings.Index(i, Space):])))
		} else {
			rpc.AddInputParameter(NewParameter(false, strings.TrimSpace(i)))
		}
	}
}

func ParseReturnArgs(values []string, rpc *Rpc) {
	returnArgs := strings.Split(values[3], Comma)
	for _, i := range returnArgs {
		if strings.HasPrefix(i, "stream") {
			rpc.AddReturnParameter(NewParameter(true, strings.TrimSpace(i[strings.Index(i, Space):])))
		} else {
			rpc.AddReturnParameter(NewParameter(true, strings.TrimSpace(i)))
		}
	}
}

func (rv *RpcVisitor) Visit(scanner Scanner, in *Line, namespace string) interface{} {
	Log.Debugf("Visiting RPC: %v\n", in)

	values := rv.RpcLineMatcher.FindStringSubmatch(in.Syntax)
	out := NewRpc(namespace, values[1], in.Comment)
	ParseInArgs(values, out)
	ParseReturnArgs(values, out)

	for scanner.Scan() {
		line := scanner.ReadLine()
		if strings.HasPrefix(line.Syntax, "option") {
			optionName := line.Syntax[strings.Index(line.Syntax, "(")+1 : strings.Index(line.Syntax, ")")]
			optionBody := ""
			for scanner.Scan() {
				oBody := scanner.ReadLine()
				optionBody += oBody.Syntax
				if line.Token == Semicolon {
					break
				}
			}
			if len(strings.TrimSpace(optionBody)) > 0 {
				out.AddRpcOption(NewRpcOption(
					Join(Period, namespace, out.Name),
					optionName,
					"",
					optionBody))
			}
		}
		if line.Token == CloseBrace {
			break
		}
	}
	return out
}
