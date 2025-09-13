// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeter


type AccessContextManagerServicePerimeterStatusIngressPoliciesIngressFrom struct {
	// Identities can be an individual user, service account, Google group, or third-party identity.
	//
	// For third-party identity, only single identities
	// are supported and other identity types are not supported.The v1 identities
	// that have the prefix user, group and serviceAccount in
	// https://cloud.google.com/iam/docs/principal-identifiers#v1 are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/access_context_manager_service_perimeter#identities AccessContextManagerServicePerimeter#identities}
	Identities *[]*string `field:"optional" json:"identities" yaml:"identities"`
	// Specifies the type of identities that are allowed access from outside the perimeter.
	//
	// If left unspecified, then members of 'identities' field will be
	// allowed access. Possible values: ["IDENTITY_TYPE_UNSPECIFIED", "ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/access_context_manager_service_perimeter#identity_type AccessContextManagerServicePerimeter#identity_type}
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
	// sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/access_context_manager_service_perimeter#sources AccessContextManagerServicePerimeter#sources}
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
}

