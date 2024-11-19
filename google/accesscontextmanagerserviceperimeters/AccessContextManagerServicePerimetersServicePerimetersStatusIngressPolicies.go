// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeters


type AccessContextManagerServicePerimetersServicePerimetersStatusIngressPolicies struct {
	// ingress_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/access_context_manager_service_perimeters#ingress_from AccessContextManagerServicePerimeters#ingress_from}
	IngressFrom *AccessContextManagerServicePerimetersServicePerimetersStatusIngressPoliciesIngressFrom `field:"optional" json:"ingressFrom" yaml:"ingressFrom"`
	// ingress_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/access_context_manager_service_perimeters#ingress_to AccessContextManagerServicePerimeters#ingress_to}
	IngressTo *AccessContextManagerServicePerimetersServicePerimetersStatusIngressPoliciesIngressTo `field:"optional" json:"ingressTo" yaml:"ingressTo"`
}

