# Proto Gen MD Diagrams

![build](https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams/actions/workflows/main.yml/badge.svg)
![coverage](https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams/actions/workflows/coverage.yml/badge.svg)
![coverage](coverage.svg)

This utility package is a compiled Go program that reads a protobuf
source directory and generates Mermaid Diagrams in <protobuf-file-name>.md files
in each directory, or the output directory with the given tree structure.

> NOTE: Only Proto 3 syntax is supported.

This utility was created to ease documentation generation of complex
Protobuf libraries to visualize models and services described in a Protocol buffers.

If you find this useful, awesome! If you find a bug, please contribute a patch,
or open a bug. Please follow the [Contributing](CONTRIBUTING.md) guidelines.

> NOTE: In order to use lcov on Apple Silicon, you'll need to install Brew and run `brew install gcovr`.

## Test Input and Output

#### Build and test 

##### Using Native Go
```shell
go build && go test ./...
./proto-gen-md-diarams -d test/protos
````

##### Using Bazel

Since Bazel is CI/CD platform, it compiles for all targets.

```shell
bazel build //... && bazel test //...

# Linux 
bazel-bin/proto-gen-md-diagrams-linux-x86_64

# OS X X64
proto-gen-md-diagrams-osx-x86_64

# OS X Apple Silicon
bazel-bin/proto-gen-md-diagrams-osx-arm64

# Windows
bazel-bin/proto-gen-md-diagrams-win-x86_64
```

| Input File                                                             | Output File                                                               |
|------------------------------------------------------------------------|---------------------------------------------------------------------------|
| [Location Protobuf](pkg/proto/data/test/location/model.proto)          | [Location Markdown](pkg/proto/data/test/location/model.proto.md)          |
| [Location Service Protobuf](pkg/proto/data/test/service/service.proto) | [Location Service Markdown](pkg/proto/data/test/service/service.proto.md) |


## Building

### Go Lang 
```shell
cd proto-gen-md-diagrams
// Build
go build && go test ./...
```

## Use and Options

```shell
./proto-gen-md-diagrams -h

Usage of ./proto-gen-md-diagrams:
  -d string
        The directoryFlag to read. (default ".")
  -debugFlag
        Enable debugging
  -o string
        Specifies the outputFlag directoryFlag, if not specified, the processor will write markdown in the proto directories. (default ".")
  -r    Read recursively. (default true)
  -v    Enable Visualization (default true)
  -w    Enable writing output (default true)

  
./proto-gen-md-diagrams -d test/protos
```

## Quick Example

### Protobuf Input

```protobuf
// A physical location that can be described with either an address
// or a set of geo coordinates.
message PhysicalLocation {
  // A postal address for the physical location.
  message Address {
    // Address type is used to identify the type of address.
    enum AddressType {
      RESIDENTIAL = 0; // A residential address
      BUSINESS = 1; // A business address
    }
    // First line of the address
    string line1 = 1;
    // Second line of the address
    string line2 = 2;
    // Third line of the address
    string line3 = 3;
    // The city or township
    string city = 4;
    // The state or province
    string state = 5;
    // The postal code
    string zipcode = 6;
    // The type of address
    AddressType type = 7;
    // Reserved for future use
    reserved 8 to 20;
  }
  // The timestamp the record was created
  google.protobuf.Timestamp created = 1;
  // The mailing address of the location
  Address address = 2;
  // Longitude degrees
  int32 longitude_degrees = 3 [json_name = 'lng_d'];
  // Longitude Minutes
  int32 longitude_minutes = 4 [json_name = 'lng_m'];
  // Longitude Seconds
  int32 longitude_seconds = 5 [json_name = 'lng_s'];
  // Longitude Degrees
  int32 latitude_degrees = 6  [json_name = 'lat_d'];
  // Latitude Minutes
  int32 latitude_minutes = 7  [json_name = 'lat_m'];
  // Latitude Seconds
  int32 latitude_seconds = 8  [json_name = 'lat_s'];
  // Latitude Direction Code
  string latitude_direction_code = 9  [json_name = 'lat_dir_code'];
  // Altitude in Meters
  double altitude_meters = 10  [json_name = 'alt_m'];
  // Additional Meta Data
  map<string, string> meta = 11;
  // Names for the location
  repeated string names = 12 [json_name = 'names'];
}
```

## Markdown Output

### Diagram
```mermaid
classDiagram
direction LR

%% A physical location that can be described with either an address or a set of geo coordinates.
class PhysicalLocation {
  + Address address
  + double altitude_meters
  + google.protobuf.Timestamp created
  + int32 latitude_degrees
  + string latitude_direction_code
  + int32 latitude_minutes
  + int32 latitude_seconds
  + int32 longitude_degrees
  + int32 longitude_minutes
  + int32 longitude_seconds
  + Map<string,  string> meta
  + List<string> names
}
PhysicalLocation --> `Address`
PhysicalLocation --> `google.protobuf.Timestamp`
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

## Description
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: test.location.PhysicalLocation</div>

A physical location that can be described with either an address or a set of geo coordinates.

| Field                   | Ordinal | Type                      | Label    | Description                          |
|-------------------------|---------|---------------------------|----------|--------------------------------------|
| address                 | 2       | Address                   |          | The mailing address of the location  |
| altitude_meters         | 10      | double                    |          | Altitude in Meters                   |
| created                 | 1       | google.protobuf.Timestamp |          | The timestamp the record was created |
| latitude_degrees        | 6       | int32                     |          | Longitude Degrees                    |
| latitude_direction_code | 9       | string                    |          | Latitude Direction Code              |
| latitude_minutes        | 7       | int32                     |          | Latitude Minutes                     |
| latitude_seconds        | 8       | int32                     |          | Latitude Seconds                     |
| longitude_degrees       | 3       | int32                     |          | Longitude degrees                    |
| longitude_minutes       | 4       | int32                     |          | Longitude Minutes                    |
| longitude_seconds       | 5       | int32                     |          | Longitude Seconds                    |
| meta                    | 11      | string, string            | Map      | Additional Meta Data                 |
| names                   | 12      | string                    | Repeated | Names for the location               |

