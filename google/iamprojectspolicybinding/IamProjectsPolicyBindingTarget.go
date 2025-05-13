// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamprojectspolicybinding


type IamProjectsPolicyBindingTarget struct {
	// Required.
	//
	// Immutable. Full Resource Name of the principal set used for principal access boundary policy bindings.
	// Examples for each one of the following supported principal set types:
	// * Project:
	//   * '//cloudresourcemanager.googleapis.com/projects/PROJECT_NUMBER'
	//   * '//cloudresourcemanager.googleapis.com/projects/PROJECT_ID'
	// * Workload Identity Pool: '//iam.googleapis.com/projects/PROJECT_NUMBER/locations/LOCATION/workloadIdentityPools/WORKLOAD_POOL_ID'
	// It must be parent by the policy binding's parent (the project).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/iam_projects_policy_binding#principal_set IamProjectsPolicyBinding#principal_set}
	PrincipalSet *string `field:"optional" json:"principalSet" yaml:"principalSet"`
}

