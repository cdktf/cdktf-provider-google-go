// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeterdryrunegresspolicy


type AccessContextManagerServicePerimeterDryRunEgressPolicyEgressToOperations struct {
	// method_selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/access_context_manager_service_perimeter_dry_run_egress_policy#method_selectors AccessContextManagerServicePerimeterDryRunEgressPolicy#method_selectors}
	MethodSelectors interface{} `field:"optional" json:"methodSelectors" yaml:"methodSelectors"`
	// The name of the API whose methods or permissions the 'IngressPolicy' or 'EgressPolicy' want to allow.
	//
	// A single 'ApiOperation' with serviceName
	// field set to '*' will allow all methods AND permissions for all services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/access_context_manager_service_perimeter_dry_run_egress_policy#service_name AccessContextManagerServicePerimeterDryRunEgressPolicy#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

