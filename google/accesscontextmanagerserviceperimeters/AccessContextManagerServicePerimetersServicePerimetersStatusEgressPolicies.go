// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeters


type AccessContextManagerServicePerimetersServicePerimetersStatusEgressPolicies struct {
	// egress_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.84.0/docs/resources/access_context_manager_service_perimeters#egress_from AccessContextManagerServicePerimeters#egress_from}
	EgressFrom *AccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressFrom `field:"optional" json:"egressFrom" yaml:"egressFrom"`
	// egress_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.84.0/docs/resources/access_context_manager_service_perimeters#egress_to AccessContextManagerServicePerimeters#egress_to}
	EgressTo *AccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressTo `field:"optional" json:"egressTo" yaml:"egressTo"`
}

