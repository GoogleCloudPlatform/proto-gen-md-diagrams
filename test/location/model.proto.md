# Package: test.location

<div class="comment"><span></span><br/><span>Copyright 2022 Google LLC</span><br/><span>Licensed under the Apache License, Version 2.0 (the "License");</span><br/><span>you may not use this file except in compliance with the License.</span><br/><span>You may obtain a copy of the License at</span><br/><span> http://www.apache.org/licenses/LICENSE-2.0</span><br/><span>Unless required by applicable law or agreed to in writing, software</span><br/><span>distributed under the License is distributed on an "AS IS" BASIS,</span><br/><span>WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.</span><br/><span>See the License for the specific language governing permissions and</span><br/><span>limitations under the License.</span><br/><span></span><br/></div>

## Imports

| Import                          | Description                          |
|---------------------------------|--------------------------------------|
| google/protobuf/timestamp.proto | Import google timestamp to identify  |



## Options

| Name                | Value                   | Description      |
|---------------------|-------------------------|------------------|
| go_package          | gcp/proto/test/location | Go Lang Options  |
| java_package        | gcp.proto.test.location | Java Options     |
| java_multiple_files | true                    |                  |



### test.location Diagram

```mermaid
classDiagram
direction LR
%% Mermaid Diagram for package: test.location

%% A physical location that can be described with either an address or a set of geo coordinates.

class PhysicalLocation {
  + google.protobuf.Timestamp created
  + Address address
  + int32 longitude_degrees
  + int32 longitude_minutes
  + int32 longitude_seconds
  + int32 latitude_degrees
  + int32 latitude_minutes
  + int32 latitude_seconds
  + string latitude_direction_code
  + double altitude_meters
  + Map~string,  string~ meta
  + List~string~ names
}
PhysicalLocation --> `Address`
PhysicalLocation --o `Address`

%% A postal address for the physical location.

class Address {
  + string line1
  + string line2
  + string line3
  + string city
  + string state
  + string zipcode
  + AddressType type
}
Address --> `AddressType`
Address --o `AddressType`
%% Address type is used to identify the type of address.

class AddressType{
  <<enumeration>>
  RESIDENTIAL
  BUSINESS
}

%% 

class PhoneNumber {
  + string country_code
  + string area_code
  + string prefix
  + string suffix
  + string extension
}

```


### PhysicalLocation Diagram

```mermaid
classDiagram
direction LR

%% A physical location that can be described with either an address or a set of geo coordinates.

class PhysicalLocation {
  + google.protobuf.Timestamp created
  + Address address
  + int32 longitude_degrees
  + int32 longitude_minutes
  + int32 longitude_seconds
  + int32 latitude_degrees
  + int32 latitude_minutes
  + int32 latitude_seconds
  + string latitude_direction_code
  + double altitude_meters
  + Map~string,  string~ meta
  + List~string~ names
}
PhysicalLocation --> `Address`
PhysicalLocation --o `Address`

%% A postal address for the physical location.

class Address {
  + string line1
  + string line2
  + string line3
  + string city
  + string state
  + string zipcode
  + AddressType type
}
Address --> `AddressType`
Address --o `AddressType`
%% Address type is used to identify the type of address.

class AddressType{
  <<enumeration>>
  RESIDENTIAL
  BUSINESS
}

```
### PhoneNumber Diagram

```mermaid
classDiagram
direction LR

%% 

class PhoneNumber {
  + string country_code
  + string area_code
  + string prefix
  + string suffix
  + string extension
}

```

## Message: PhysicalLocation
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: test.location.PhysicalLocation</div>

<div class="comment"><span>A physical location that can be described with either an address or a set of geo coordinates.</span><br/></div>

| Field                   | Ordinal | Type                      | Label    | Description                           |
|-------------------------|---------|---------------------------|----------|---------------------------------------|
| created                 | 1       | google.protobuf.Timestamp |          | The timestamp the record was created  |
| address                 | 2       | Address                   |          | The mailing address of the location   |
| longitude_degrees       | 3       | int32                     |          | Longitude degrees                     |
| longitude_minutes       | 4       | int32                     |          | Longitude Minutes                     |
| longitude_seconds       | 5       | int32                     |          | Longitude Seconds                     |
| latitude_degrees        | 6       | int32                     |          | Longitude Degrees                     |
| latitude_minutes        | 7       | int32                     |          | Latitude Minutes                      |
| latitude_seconds        | 8       | int32                     |          | Latitude Seconds                      |
| latitude_direction_code | 9       | string                    |          | Latitude Direction Code               |
| altitude_meters         | 10      | double                    |          | Altitude in Meters                    |
| meta                    | 11      | string, string            | Map      | Additional Meta Data                  |
| names                   | 12      | string                    | Repeated | Names for the location                |



### Address Diagram

```mermaid
classDiagram
direction LR

%% A postal address for the physical location.

class Address {
  + string line1
  + string line2
  + string line3
  + string city
  + string state
  + string zipcode
  + AddressType type
}
Address --> `AddressType`
Address --o `AddressType`
%% Address type is used to identify the type of address.

class AddressType{
  <<enumeration>>
  RESIDENTIAL
  BUSINESS
}

```
### AddressType Diagram

```mermaid
classDiagram
direction LR
%% Address type is used to identify the type of address.

class AddressType{
  <<enumeration>>
  RESIDENTIAL
  BUSINESS
}

```

## Message: Address
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: test.location.PhysicalLocation.Address</div>

<div class="comment"><span>A postal address for the physical location.</span><br/></div>

| Field   | Ordinal | Type        | Label | Description                 |
|---------|---------|-------------|-------|-----------------------------|
| line1   | 1       | string      |       | First line of the address   |
| line2   | 2       | string      |       | Second line of the address  |
| line3   | 3       | string      |       | Third line of the address   |
| city    | 4       | string      |       | The city or township        |
| state   | 5       | string      |       | The state or province       |
| zipcode | 6       | string      |       | The postal code             |
| type    | 7       | AddressType |       | The type of address         |


## Enum: AddressType
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: test.location.PhysicalLocation.Address.AddressType</div>

<div class="comment"><span>Address type is used to identify the type of address.</span><br/></div>

| Name        | Ordinal | Description            |
|-------------|---------|------------------------|
| RESIDENTIAL | 0       | A residential address  |
| BUSINESS    | 1       | A business address     |




## Message: PhoneNumber
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: test.location.PhoneNumber</div>

<div class="comment"><span></span><br/></div>

| Field        | Ordinal | Type   | Label | Description |
|--------------|---------|--------|-------|-------------|
| country_code | 1       | string |       |             |
| area_code    | 2       | string |       |             |
| prefix       | 3       | string |       |             |
| suffix       | 4       | string |       |             |
| extension    | 5       | string |       |             |






<!-- Created by: Proto Diagram Tool -->
<!-- https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams -->
