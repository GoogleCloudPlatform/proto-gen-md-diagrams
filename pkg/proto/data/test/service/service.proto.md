# Package: test.service

<div class="comment"><span></span><br/><span>Copyright 2022 Google LLC</span><br/><span>Licensed under the Apache License, Version 2.0 (the "License");</span><br/><span>you may not use this file except in compliance with the License.</span><br/><span>You may obtain a copy of the License at</span><br/><span> http://www.apache.org/licenses/LICENSE-2.0</span><br/><span>Unless required by applicable law or agreed to in writing, software</span><br/><span>distributed under the License is distributed on an "AS IS" BASIS,</span><br/><span>WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.</span><br/><span>See the License for the specific language governing permissions and</span><br/><span>limitations under the License.</span><br/><span></span><br/></div>

## Imports

| Import                       | Description |
|------------------------------|-------------|
| test/location/model.proto    |             |
| google/protobuf/empty.proto  |             |
| google/api/annotations.proto |             |



## Options

| Name                | Value                                       | Description      |
|---------------------|---------------------------------------------|------------------|
| go_package          | github.com/rrmcguinness/proto/test/location | Go Lang Options  |
| java_package        | com.github.rrmcguinness.proto.test.location | Java Options     |
| java_multiple_files | true                                        |                  |



## Service: LocationService
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: test.service</div>

<div class="comment"><span></span><br/><span>The LocationService is responsible for CRUD operations of Physical Locations.</span><br/><span></span><br/></div>

### LocationService Diagram

```mermaid
classDiagram
direction LR
class LocationService {
  <<service>>
  +List(Empty) Stream~PhysicalLocation~
}
LocationService --> `google.protobuf.Empty`
LocationService --o `test.location.PhysicalLocation`

```

| Method | Parameter (In) | Parameter (Out)            | Description                                |
|--------|----------------|----------------------------|--------------------------------------------|
| List   | Empty          | Stream\<PhysicalLocation\> | List returns a list of physical locations  |







<!-- Created by: Proto Diagram Tool -->
<!-- https://github.com/rrmcguinness/proto-diagram-tool -->

