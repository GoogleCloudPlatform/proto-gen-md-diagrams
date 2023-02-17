module github.com/GoogleCloudPlatform/proto-gen-md-diagrams

go 1.19

replace github.com/GoogleCloudPlatform/proto-gen-md-diagrams => ./cmd

replace github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/api => ./pkg/api

replace github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/pb => ./pkg/pb

replace github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/logging => ./pkg/logging

replace github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/reader => ./pkg/reader

require github.com/stretchr/testify v1.8.1

require github.com/davecgh/go-spew v1.1.1 // indirect

require github.com/pmezard/go-difflib v1.0.0 // indirect

require gopkg.in/yaml.v3 v3.0.1 // indirect
