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

type Parameter struct {
	Stream bool
	Type   string
}

func NewParameter(stream bool, t string) *Parameter {
	return &Parameter{Stream: stream, Type: t}
}

type RpcOption struct {
	*Qualified
	Body string
}

func NewRpcOption(namespace string, name string, comment Comment, body string) *RpcOption {
	return &RpcOption{
		Qualified: &Qualified{
			Qualifier: namespace,
			Name:      name,
			Comment:   comment,
		},
		Body: body,
	}
}

type Rpc struct {
	*Qualified
	InputParameters  []*Parameter
	ReturnParameters []*Parameter
	Options          []*RpcOption
}

func NewRpc(namespace string, name string, comment Comment) *Rpc {
	return &Rpc{
		Qualified: &Qualified{
			Qualifier: namespace,
			Name:      name,
			Comment:   comment,
		},
		InputParameters:  make([]*Parameter, 0),
		ReturnParameters: make([]*Parameter, 0),
		Options:          make([]*RpcOption, 0),
	}
}

func (rpc *Rpc) AddInputParameter(params ...*Parameter) {
	rpc.InputParameters = append(rpc.InputParameters, params...)
}

func (rpc *Rpc) AddReturnParameter(params ...*Parameter) {
	rpc.ReturnParameters = append(rpc.ReturnParameters, params...)
}

func (rpc *Rpc) AddRpcOption(options ...*RpcOption) {
	rpc.Options = append(rpc.Options, options...)
}
