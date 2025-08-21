// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccorganizationcustommodule


type SccOrganizationCustomModuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/scc_organization_custom_module#create SccOrganizationCustomModule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/scc_organization_custom_module#delete SccOrganizationCustomModule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/scc_organization_custom_module#update SccOrganizationCustomModule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

