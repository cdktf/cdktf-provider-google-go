// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamprincipalaccessboundarypolicy


type IamPrincipalAccessBoundaryPolicyDetailsRules struct {
	// The access relationship of principals to the resources in this rule. Possible values: ALLOW.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/iam_principal_access_boundary_policy#effect IamPrincipalAccessBoundaryPolicy#effect}
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// A list of Cloud Resource Manager resources.
	//
	// The resource
	// and all the descendants are included. The number of resources in a policy
	// is limited to 500 across all rules.
	// The following resource types are supported:
	// * Organizations, such as '//cloudresourcemanager.googleapis.com/organizations/123'.
	// * Folders, such as '//cloudresourcemanager.googleapis.com/folders/123'.
	// * Projects, such as '//cloudresourcemanager.googleapis.com/projects/123'
	// or '//cloudresourcemanager.googleapis.com/projects/my-project-id'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/iam_principal_access_boundary_policy#resources IamPrincipalAccessBoundaryPolicy#resources}
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
	// The description of the principal access boundary policy rule. Must be less than or equal to 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/iam_principal_access_boundary_policy#description IamPrincipalAccessBoundaryPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

