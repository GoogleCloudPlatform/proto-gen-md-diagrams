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

type service struct {
	api.Qualified
	remoteProcedureCalls []api.RPC
}

func (s *service) RemoteProcedureCalls() []api.RPC {
	return s.remoteProcedureCalls
}

func (s *service) AddRPC(rpc api.RPC) api.Service {
	s.remoteProcedureCalls = append(s.remoteProcedureCalls, rpc)
	return s
}

// rPC Parameters

type rPCParameter struct {
	stream bool
	kind   string
}

func (r rPCParameter) Stream() bool {
	return r.stream
}

func (r rPCParameter) Kind() string {
	return r.kind
}

// rPC Options

type rPCOption struct {
	api.Qualified
	body string
}

func (r *rPCOption) Body() string {
	return r.body
}

type rPC struct {
	api.Qualified
	inputParameters  []api.RPCParameter
	returnParameters []api.RPCParameter
	rpcOptions       []api.RPCOption
}

func (r *rPC) InputParameters() []api.RPCParameter {
	return r.inputParameters
}

func (r *rPC) AddInputParameter(stream bool, kind string) api.RPC {
	r.inputParameters = append(r.inputParameters, &rPCParameter{stream: stream, kind: kind})
	return r
}

func (r *rPC) ReturnParameters() []api.RPCParameter {
	return r.returnParameters
}

func (r *rPC) AddReturnParameter(stream bool, kind string) api.RPC {
	r.returnParameters = append(r.returnParameters, &rPCParameter{stream: stream, kind: kind})
	return r
}

func (r *rPC) RPCOptions() []api.RPCOption {
	return r.rpcOptions
}

func (r *rPC) AddRPCOption(name string, comment string, body string) api.RPC {
	r.rpcOptions = append(r.rpcOptions, &rPCOption{
		Qualified: newQualified(name, name, comment),
		body:      body,
	})
	return r
}
