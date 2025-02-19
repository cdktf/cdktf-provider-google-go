// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamorganizationspolicybinding


type IamOrganizationsPolicyBindingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/iam_organizations_policy_binding#create IamOrganizationsPolicyBinding#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/iam_organizations_policy_binding#delete IamOrganizationsPolicyBinding#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/iam_organizations_policy_binding#update IamOrganizationsPolicyBinding#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

