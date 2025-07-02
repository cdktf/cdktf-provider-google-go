// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeterdryrunegresspolicy


type AccessContextManagerServicePerimeterDryRunEgressPolicyEgressFrom struct {
	// Identities can be an individual user, service account, Google group, or third-party identity.
	//
	// For third-party identity, only single identities
	// are supported and other identity types are not supported.The v1 identities
	// that have the prefix user, group and serviceAccount in
	// https://cloud.google.com/iam/docs/principal-identifiers#v1 are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/access_context_manager_service_perimeter_dry_run_egress_policy#identities AccessContextManagerServicePerimeterDryRunEgressPolicy#identities}
	Identities *[]*string `field:"optional" json:"identities" yaml:"identities"`
	// Specifies the type of identities that are allowed access to outside the perimeter.
	//
	// If left unspecified, then members of 'identities' field will
	// be allowed access. Possible values: ["ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/access_context_manager_service_perimeter_dry_run_egress_policy#identity_type AccessContextManagerServicePerimeterDryRunEgressPolicy#identity_type}
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
	// Whether to enforce traffic restrictions based on 'sources' field.
	//
	// If the 'sources' field is non-empty, then this field must be set to 'SOURCE_RESTRICTION_ENABLED'. Possible values: ["SOURCE_RESTRICTION_ENABLED", "SOURCE_RESTRICTION_DISABLED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/access_context_manager_service_perimeter_dry_run_egress_policy#source_restriction AccessContextManagerServicePerimeterDryRunEgressPolicy#source_restriction}
	SourceRestriction *string `field:"optional" json:"sourceRestriction" yaml:"sourceRestriction"`
	// sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/access_context_manager_service_perimeter_dry_run_egress_policy#sources AccessContextManagerServicePerimeterDryRunEgressPolicy#sources}
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
}

