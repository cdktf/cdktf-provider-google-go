// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigpatchdeployment


type OsConfigPatchDeploymentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/os_config_patch_deployment#create OsConfigPatchDeployment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/os_config_patch_deployment#delete OsConfigPatchDeployment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

