// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanageraccesslevel


type AccessContextManagerAccessLevelTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/access_context_manager_access_level#create AccessContextManagerAccessLevel#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/access_context_manager_access_level#delete AccessContextManagerAccessLevel#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/access_context_manager_access_level#update AccessContextManagerAccessLevel#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

