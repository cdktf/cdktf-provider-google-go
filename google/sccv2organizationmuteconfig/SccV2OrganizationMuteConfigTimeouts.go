// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccv2organizationmuteconfig


type SccV2OrganizationMuteConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/scc_v2_organization_mute_config#create SccV2OrganizationMuteConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/scc_v2_organization_mute_config#delete SccV2OrganizationMuteConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/scc_v2_organization_mute_config#update SccV2OrganizationMuteConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

