// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanageraccesslevels


type AccessContextManagerAccessLevelsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/access_context_manager_access_levels#create AccessContextManagerAccessLevels#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/access_context_manager_access_levels#delete AccessContextManagerAccessLevels#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/access_context_manager_access_levels#update AccessContextManagerAccessLevels#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

