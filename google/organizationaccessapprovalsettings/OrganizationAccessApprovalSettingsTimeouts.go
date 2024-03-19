// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationaccessapprovalsettings


type OrganizationAccessApprovalSettingsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/organization_access_approval_settings#create OrganizationAccessApprovalSettings#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/organization_access_approval_settings#delete OrganizationAccessApprovalSettings#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/organization_access_approval_settings#update OrganizationAccessApprovalSettings#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

