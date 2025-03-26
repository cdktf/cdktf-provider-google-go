// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamprojectspolicybinding


type IamProjectsPolicyBindingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/iam_projects_policy_binding#create IamProjectsPolicyBinding#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/iam_projects_policy_binding#delete IamProjectsPolicyBinding#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/iam_projects_policy_binding#update IamProjectsPolicyBinding#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

