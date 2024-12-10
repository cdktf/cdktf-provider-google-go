// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamprojectspolicybinding


type IamProjectsPolicyBindingTarget struct {
	// Required.
	//
	// Immutable. The resource name of the policy to be bound.
	// The binding parent and policy must belong to the same Organization (or Project).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/iam_projects_policy_binding#principal_set IamProjectsPolicyBinding#principal_set}
	PrincipalSet *string `field:"optional" json:"principalSet" yaml:"principalSet"`
}

