// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeter


type AccessContextManagerServicePerimeterStatusEgressPolicies struct {
	// egress_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/access_context_manager_service_perimeter#egress_from AccessContextManagerServicePerimeter#egress_from}
	EgressFrom *AccessContextManagerServicePerimeterStatusEgressPoliciesEgressFrom `field:"optional" json:"egressFrom" yaml:"egressFrom"`
	// egress_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/access_context_manager_service_perimeter#egress_to AccessContextManagerServicePerimeter#egress_to}
	EgressTo *AccessContextManagerServicePerimeterStatusEgressPoliciesEgressTo `field:"optional" json:"egressTo" yaml:"egressTo"`
	// Human readable title. Must be unique within the perimeter. Does not affect behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/access_context_manager_service_perimeter#title AccessContextManagerServicePerimeter#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

