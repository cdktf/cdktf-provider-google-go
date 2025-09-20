// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccv2foldermuteconfig


type SccV2FolderMuteConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/scc_v2_folder_mute_config#create SccV2FolderMuteConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/scc_v2_folder_mute_config#delete SccV2FolderMuteConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/scc_v2_folder_mute_config#update SccV2FolderMuteConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

