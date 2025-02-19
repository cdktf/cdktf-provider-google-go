// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamfolderspolicybinding


type IamFoldersPolicyBindingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/iam_folders_policy_binding#create IamFoldersPolicyBinding#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/iam_folders_policy_binding#delete IamFoldersPolicyBinding#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/iam_folders_policy_binding#update IamFoldersPolicyBinding#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

