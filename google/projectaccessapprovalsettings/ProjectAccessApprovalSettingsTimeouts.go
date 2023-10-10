// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectaccessapprovalsettings


type ProjectAccessApprovalSettingsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.1.0/docs/resources/project_access_approval_settings#create ProjectAccessApprovalSettings#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.1.0/docs/resources/project_access_approval_settings#delete ProjectAccessApprovalSettings#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.1.0/docs/resources/project_access_approval_settings#update ProjectAccessApprovalSettings#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

