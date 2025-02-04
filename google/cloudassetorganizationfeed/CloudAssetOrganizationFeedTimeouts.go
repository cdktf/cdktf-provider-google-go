// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudassetorganizationfeed


type CloudAssetOrganizationFeedTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/cloud_asset_organization_feed#create CloudAssetOrganizationFeed#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/cloud_asset_organization_feed#delete CloudAssetOrganizationFeed#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/cloud_asset_organization_feed#update CloudAssetOrganizationFeed#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

