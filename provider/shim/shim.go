package shim

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"
	internalprovider "github.com/wundergraph/cosmo/terraform-provider-cosmo/internal/provider"
)

func NewProvider(version string) func() provider.Provider {
	return internalprovider.New(version)
}
