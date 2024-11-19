// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeteringresspolicy


type AccessContextManagerServicePerimeterIngressPolicyIngressFrom struct {
	// Identities can be an individual user, service account, Google group, or third-party identity.
	//
	// For third-party identity, only single identities
	// are supported and other identity types are not supported.The v1 identities
	// that have the prefix user, group and serviceAccount in
	// https://cloud.google.com/iam/docs/principal-identifiers#v1 are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/access_context_manager_service_perimeter_ingress_policy#identities AccessContextManagerServicePerimeterIngressPolicy#identities}
	Identities *[]*string `field:"optional" json:"identities" yaml:"identities"`
	// Specifies the type of identities that are allowed access from outside the perimeter.
	//
	// If left unspecified, then members of 'identities' field will be
	// allowed access. Possible values: ["ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/access_context_manager_service_perimeter_ingress_policy#identity_type AccessContextManagerServicePerimeterIngressPolicy#identity_type}
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
	// sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/access_context_manager_service_perimeter_ingress_policy#sources AccessContextManagerServicePerimeterIngressPolicy#sources}
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
}

