// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeterdryruningresspolicy


type AccessContextManagerServicePerimeterDryRunIngressPolicyIngressToOperations struct {
	// method_selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/access_context_manager_service_perimeter_dry_run_ingress_policy#method_selectors AccessContextManagerServicePerimeterDryRunIngressPolicy#method_selectors}
	MethodSelectors interface{} `field:"optional" json:"methodSelectors" yaml:"methodSelectors"`
	// The name of the API whose methods or permissions the 'IngressPolicy' or 'EgressPolicy' want to allow.
	//
	// A single 'ApiOperation' with 'serviceName'
	// field set to '*' will allow all methods AND permissions for all services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/access_context_manager_service_perimeter_dry_run_ingress_policy#service_name AccessContextManagerServicePerimeterDryRunIngressPolicy#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

