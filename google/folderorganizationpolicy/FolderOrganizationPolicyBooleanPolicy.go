// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package folderorganizationpolicy


type FolderOrganizationPolicyBooleanPolicy struct {
	// If true, then the Policy is enforced. If false, then any configuration is acceptable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/folder_organization_policy#enforced FolderOrganizationPolicy#enforced}
	Enforced interface{} `field:"required" json:"enforced" yaml:"enforced"`
}

