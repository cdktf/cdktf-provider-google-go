// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccv2organizationsource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SccV2OrganizationSourceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The source’s display name.
	//
	// A source’s display name must be unique
	// amongst its siblings, for example, two sources with the same parent
	// can't share the same display name. The display name must start and end
	// with a letter or digit, may contain letters, digits, spaces, hyphens,
	// and underscores, and can be no longer than 32 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/scc_v2_organization_source#display_name SccV2OrganizationSource#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The organization whose Cloud Security Command Center the Source lives in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/scc_v2_organization_source#organization SccV2OrganizationSource#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The description of the source (max of 1024 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/scc_v2_organization_source#description SccV2OrganizationSource#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/scc_v2_organization_source#id SccV2OrganizationSource#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/scc_v2_organization_source#timeouts SccV2OrganizationSource#timeouts}
	Timeouts *SccV2OrganizationSourceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

