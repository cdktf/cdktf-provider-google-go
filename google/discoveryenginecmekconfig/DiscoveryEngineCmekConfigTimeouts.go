// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginecmekconfig


type DiscoveryEngineCmekConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/discovery_engine_cmek_config#create DiscoveryEngineCmekConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/discovery_engine_cmek_config#delete DiscoveryEngineCmekConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/discovery_engine_cmek_config#update DiscoveryEngineCmekConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

