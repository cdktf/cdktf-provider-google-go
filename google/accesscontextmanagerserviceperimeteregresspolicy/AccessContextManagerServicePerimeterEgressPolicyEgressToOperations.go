// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeteregresspolicy


type AccessContextManagerServicePerimeterEgressPolicyEgressToOperations struct {
	// method_selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/access_context_manager_service_perimeter_egress_policy#method_selectors AccessContextManagerServicePerimeterEgressPolicy#method_selectors}
	MethodSelectors interface{} `field:"optional" json:"methodSelectors" yaml:"methodSelectors"`
	// The name of the API whose methods or permissions the 'IngressPolicy' or 'EgressPolicy' want to allow.
	//
	// A single 'ApiOperation' with serviceName
	// field set to '*' will allow all methods AND permissions for all services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/access_context_manager_service_perimeter_egress_policy#service_name AccessContextManagerServicePerimeterEgressPolicy#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

