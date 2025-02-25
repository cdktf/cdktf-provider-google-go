// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginetargetsite


type DiscoveryEngineTargetSiteTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/discovery_engine_target_site#create DiscoveryEngineTargetSite#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/discovery_engine_target_site#delete DiscoveryEngineTargetSite#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

