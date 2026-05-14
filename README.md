# cosmo Resource Provider

The cosmo Resource Provider lets you manage [cosmo](https://github.com/wundergraph/terraform-provider-cosmo) resources.

## Development note: Terraform internal provider and shim pattern

This repository uses the shim pattern recommended by Pulumi's Terraform bridge boilerplate.

Why this is needed:

- The upstream Terraform provider exposes its implementation under an `internal/...` package.
- Go does not allow importing `internal` packages from outside the parent module path.
- Importing `github.com/wundergraph/cosmo/terraform-provider-cosmo/internal/provider` directly from this Pulumi provider would fail with `use of internal package ... not allowed`.

Chosen solution:

- Keep the direct `internal/provider` import only inside `provider/shim/shim.go`, in a module that lives under the same upstream module path (`github.com/wundergraph/cosmo/terraform-provider-cosmo/shim`).
- In `provider/resources.go`, import the shim module and initialize the provider via `pfbridge.ShimProvider(shim.NewProvider(version.Version)())`.
- In `provider/go.mod`, use local `replace` directives during development:
	- `github.com/wundergraph/cosmo/terraform-provider-cosmo => ../upstream`
	- `github.com/wundergraph/cosmo/terraform-provider-cosmo/shim => ./shim`

This keeps the bridge code clean, respects Go `internal` visibility rules, and aligns with Pulumi's documented approach for Terraform providers that hide constructors in `internal` packages.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumi/cosmo
```

or `yarn`:

```bash
yarn add @pulumi/cosmo
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_cosmo
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumi/pulumi-cosmo/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.cosmo
```
