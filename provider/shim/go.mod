module github.com/wundergraph/cosmo/terraform-provider-cosmo/shim

go 1.22.0

require (
	github.com/hashicorp/terraform-plugin-framework v1.11.0
	github.com/wundergraph/cosmo/terraform-provider-cosmo v0.0.0
)

require (
	connectrpc.com/connect v1.16.2 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/terraform-plugin-framework-validators v0.13.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.23.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.9.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/wundergraph/cosmo/connect-go v0.0.0-20241203152720-979e5a780c8e // indirect
	golang.org/x/sys v0.23.0 // indirect
	google.golang.org/protobuf v1.34.0 // indirect
)

replace github.com/wundergraph/cosmo/terraform-provider-cosmo => ../../upstream
