module github.com/cloudevents/sdk-go/sql/v2

go 1.18

require (
	github.com/antlr/antlr4/runtime/Go/antlr v1.4.10
	github.com/cloudevents/sdk-go v1.2.0
	github.com/cloudevents/sdk-go/v2 v2.5.0
	github.com/stretchr/testify v1.8.0
	sigs.k8s.io/yaml v1.3.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/lightstep/tracecontext.go v0.0.0-20181129014701-1757c391b1ac // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opencensus.io v0.22.0 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/cloudevents/sdk-go/v2 => ../../v2
replace github.com/cloudevents/sdk-go/sql/v2 => github.com/dgeorgievski/sdk-go/sql/v2
