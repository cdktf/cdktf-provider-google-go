// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package folderorganizationpolicy


type FolderOrganizationPolicyRestorePolicy struct {
	// May only be set to true. If set, then the default Policy is restored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/folder_organization_policy#default FolderOrganizationPolicy#default}
	Default interface{} `field:"required" json:"default" yaml:"default"`
}

