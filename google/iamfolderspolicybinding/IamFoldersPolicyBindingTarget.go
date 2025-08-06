// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamfolderspolicybinding


type IamFoldersPolicyBindingTarget struct {
	// Required.
	//
	// Immutable. Full Resource Name of the principal set used for principal access boundary policy bindings.
	// Examples for each one of the following supported principal set types:
	// * Folder: '//cloudresourcemanager.googleapis.com/folders/FOLDER_ID'
	// It must be parent by the policy binding's parent (the folder).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/iam_folders_policy_binding#principal_set IamFoldersPolicyBinding#principal_set}
	PrincipalSet *string `field:"optional" json:"principalSet" yaml:"principalSet"`
}

