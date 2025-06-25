// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccv2projectmuteconfig


type SccV2ProjectMuteConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/scc_v2_project_mute_config#create SccV2ProjectMuteConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/scc_v2_project_mute_config#delete SccV2ProjectMuteConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/scc_v2_project_mute_config#update SccV2ProjectMuteConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

