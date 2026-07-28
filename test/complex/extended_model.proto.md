# Package: test.complex

<div class="comment"><span></span><br/></div>

## Imports

| Import                          | Description |
|---------------------------------|-------------|
| google/protobuf/timestamp.proto |             |



## Options

| Name         | Value            | Description |
|--------------|------------------|-------------|
| java_package | com.test.complex |             |



### test.complex Diagram

```mermaid
classDiagram
direction LR
%% Mermaid Diagram for package: test.complex

%% 

class CustomData {
  + Optional~string~ id
}

%% 

class ComplexMessage {
  + Optional~google.protobuf.Timestamp~ created
  + Oneof~string~ text [oneof: payload]
  + Oneof~CustomData~ data [oneof: payload]
  + Optional~string~ old_field [deprecated]
  + Optional~Result~ result
}
ComplexMessage --> `CustomData`
ComplexMessage --> `Result`
ComplexMessage --o `Result`

%% 

class Result {
  + Optional~string~ status
}
class ComplexService {
  <<service>>
  +ExecuteUnary (CustomData) ComplexMessage
  +ExecuteClientStream (Stream~CustomData~) ComplexMessage
  +ExecuteServerStream (CustomData) Stream~ComplexMessage~
  +ExecuteBidiStream (Stream~CustomData~) Stream~ComplexMessage~
}
ComplexService --> `CustomData`
ComplexService --> `ComplexMessage`
ComplexService --o `CustomData` : client stream
ComplexService --> `ComplexMessage`
ComplexService --> `CustomData`
ComplexService --o `ComplexMessage` : server stream
ComplexService --o `CustomData` : client stream
ComplexService --o `ComplexMessage` : server stream

```

## Service: ComplexService
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: test.complex</div>

<div class="comment"><span></span><br/></div>

### ComplexService Diagram

```mermaid
classDiagram
direction LR
class ComplexService {
  <<service>>
  +ExecuteUnary (CustomData) ComplexMessage
  +ExecuteClientStream (Stream~CustomData~) ComplexMessage
  +ExecuteServerStream (CustomData) Stream~ComplexMessage~
  +ExecuteBidiStream (Stream~CustomData~) Stream~ComplexMessage~
}
ComplexService --> `CustomData`
ComplexService --> `ComplexMessage`
ComplexService --o `CustomData` : client stream
ComplexService --> `ComplexMessage`
ComplexService --> `CustomData`
ComplexService --o `ComplexMessage` : server stream
ComplexService --o `CustomData` : client stream
ComplexService --o `ComplexMessage` : server stream

```

| Method               | Parameter (In)       | Parameter (Out)          | Description |
|----------------------|----------------------|--------------------------|-------------|
| ExecuteUnary         | CustomData           | ComplexMessage           |             |
| ExecuteClientStream  | Stream\<CustomData\> | ComplexMessage           |             |
| ExecuteServerStream  | CustomData           | Stream\<ComplexMessage\> |             |
| ExecuteBidiStream    | Stream\<CustomData\> | Stream\<ComplexMessage\> |             |



### CustomData Diagram

```mermaid
classDiagram
direction LR

%% 

class CustomData {
  + Optional~string~ id
}

```
### ComplexMessage Diagram

```mermaid
classDiagram
direction LR

%% 

class ComplexMessage {
  + Optional~google.protobuf.Timestamp~ created
  + Oneof~string~ text [oneof: payload]
  + Oneof~CustomData~ data [oneof: payload]
  + Optional~string~ old_field [deprecated]
  + Optional~Result~ result
}
ComplexMessage --> `CustomData`
ComplexMessage --> `Result`
ComplexMessage --o `Result`

%% 

class Result {
  + Optional~string~ status
}

```

## Message: CustomData
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: test.complex.CustomData</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type   | Label    | Description |
|-------|---------|--------|----------|-------------|
| id    | 1       | string | Optional |             |




## Message: ComplexMessage
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: test.complex.ComplexMessage</div>

<div class="comment"><span></span><br/></div>

| Field     | Ordinal | Type                      | Label                | Description |
|-----------|---------|---------------------------|----------------------|-------------|
| created   | 1       | google.protobuf.Timestamp | Optional             |             |
| text      | 2       | string                    | Oneof (payload)      |             |
| data      | 3       | CustomData                | Oneof (payload)      |             |
| old_field | 4       | string                    | Optional, Deprecated |             |
| result    | 5       | Result                    | Optional             |             |



### Result Diagram

```mermaid
classDiagram
direction LR

%% 

class Result {
  + Optional~string~ status
}

```

## Message: Result
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: test.complex.ComplexMessage.Result</div>

<div class="comment"><span></span><br/></div>

| Field  | Ordinal | Type   | Label    | Description |
|--------|---------|--------|----------|-------------|
| status | 1       | string | Optional |             |






<!-- Created by: Proto Diagram Tool -->
<!-- https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams -->
