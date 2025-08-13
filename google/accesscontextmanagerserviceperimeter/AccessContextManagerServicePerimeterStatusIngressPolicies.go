// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeter


type AccessContextManagerServicePerimeterStatusIngressPolicies struct {
	// ingress_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/access_context_manager_service_perimeter#ingress_from AccessContextManagerServicePerimeter#ingress_from}
	IngressFrom *AccessContextManagerServicePerimeterStatusIngressPoliciesIngressFrom `field:"optional" json:"ingressFrom" yaml:"ingressFrom"`
	// ingress_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/access_context_manager_service_perimeter#ingress_to AccessContextManagerServicePerimeter#ingress_to}
	IngressTo *AccessContextManagerServicePerimeterStatusIngressPoliciesIngressTo `field:"optional" json:"ingressTo" yaml:"ingressTo"`
	// Human readable title. Must be unique within the perimeter. Does not affect behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/access_context_manager_service_perimeter#title AccessContextManagerServicePerimeter#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

