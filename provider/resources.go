// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cosmo

import (
	"path"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	shim "github.com/wundergraph/cosmo/terraform-provider-cosmo/shim"
	"github.com/wundergraph/pulumi-cosmo/provider/pkg/version"

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "cosmo"
	// modules:
	mainMod = "index" // the cosmo module
)

//go:embed cmd/pulumi-resource-cosmo/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		// Use pfbridge for terraform-plugin-framework based providers
		P: pfbridge.ShimProvider(shim.NewProvider(version.Version)()),

		Name:              "cosmo",
		Version:           version.Version,
		DisplayName:       "Cosmo",
		Publisher:         "WunderGraph",
		LogoURL:           "",
		PluginDownloadURL: "https://github.com/wundergraph/pulumi-cosmo/releases/download/v${VERSION}/",
		Description:       "A Pulumi package for creating and managing WunderGraph Cosmo resources.",
		Keywords:          []string{"cosmo", "wundergraph", "category/cloud"},
		License:           "Apache-2.0",
		Homepage:          "https://cosmo-docs.wundergraph.com/",
		Repository:        "https://github.com/wundergraph/pulumi-cosmo",
		GitHubOrg:         "wundergraph",
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
		Config:            map[string]*tfbridge.SchemaInfo{},
		JavaScript: &tfbridge.JavaScriptInfo{
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			RespectSchemaVersion: true,
			PyProject:            struct{ Enabled bool }{true},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				"github.com/wundergraph/pulumi-cosmo/sdk/",
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			GenerateExtraInputTypes:        true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.MustComputeTokens(tokens.SingleModule("cosmo_", mainMod,
		tokens.MakeStandard(mainPkg)))

	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
