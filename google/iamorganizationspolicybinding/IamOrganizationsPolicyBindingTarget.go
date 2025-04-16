// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamorganizationspolicybinding


type IamOrganizationsPolicyBindingTarget struct {
	// Required.
	//
	// Immutable. Full Resource Name of the principal set used for principal access boundary policy bindings.
	// Examples for each one of the following supported principal set types:
	// * Organization '//cloudresourcemanager.googleapis.com/organizations/ORGANIZATION_ID'
	// * Workforce Identity: '//iam.googleapis.com/locations/global/workforcePools/WORKFORCE_POOL_ID'
	// * Workspace Identity: '//iam.googleapis.com/locations/global/workspace/WORKSPACE_ID'
	// It must be parent by the policy binding's parent (the organization).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/iam_organizations_policy_binding#principal_set IamOrganizationsPolicyBinding#principal_set}
	PrincipalSet *string `field:"optional" json:"principalSet" yaml:"principalSet"`
}

