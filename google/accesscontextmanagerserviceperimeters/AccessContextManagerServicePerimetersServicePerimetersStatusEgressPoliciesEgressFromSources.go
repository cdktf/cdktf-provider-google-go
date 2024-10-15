// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeters


type AccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressFromSources struct {
	// An AccessLevel resource name that allows resources outside the ServicePerimeter to be accessed from the inside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/access_context_manager_service_perimeters#access_level AccessContextManagerServicePerimeters#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
}

