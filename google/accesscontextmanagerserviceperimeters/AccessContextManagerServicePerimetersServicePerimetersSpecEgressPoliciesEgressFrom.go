// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeters


type AccessContextManagerServicePerimetersServicePerimetersSpecEgressPoliciesEgressFrom struct {
	// 'A list of identities that are allowed access through this 'EgressPolicy'.
	//
	// To specify an identity or identity group, use the IAM v1 format
	// specified [here](https://cloud.google.com/iam/docs/principal-identifiers.md#v1).
	// The following prefixes are supprted: user, group, serviceAccount, principal, and principalSet.'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/access_context_manager_service_perimeters#identities AccessContextManagerServicePerimeters#identities}
	Identities *[]*string `field:"optional" json:"identities" yaml:"identities"`
	// Specifies the type of identities that are allowed access to outside the perimeter.
	//
	// If left unspecified, then members of 'identities' field will
	// be allowed access. Possible values: ["IDENTITY_TYPE_UNSPECIFIED", "ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/access_context_manager_service_perimeters#identity_type AccessContextManagerServicePerimeters#identity_type}
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
	// Whether to enforce traffic restrictions based on 'sources' field.
	//
	// If the 'sources' field is non-empty, then this field must be set to 'SOURCE_RESTRICTION_ENABLED'. Possible values: ["SOURCE_RESTRICTION_UNSPECIFIED", "SOURCE_RESTRICTION_ENABLED", "SOURCE_RESTRICTION_DISABLED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/access_context_manager_service_perimeters#source_restriction AccessContextManagerServicePerimeters#source_restriction}
	SourceRestriction *string `field:"optional" json:"sourceRestriction" yaml:"sourceRestriction"`
	// sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/access_context_manager_service_perimeters#sources AccessContextManagerServicePerimeters#sources}
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
}

