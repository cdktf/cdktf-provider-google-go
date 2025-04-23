// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeter


type AccessContextManagerServicePerimeterStatusIngressPoliciesIngressTo struct {
	// operations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/access_context_manager_service_perimeter#operations AccessContextManagerServicePerimeter#operations}
	Operations interface{} `field:"optional" json:"operations" yaml:"operations"`
	// A list of resources, currently only projects in the form 'projects/<projectnumber>', protected by this 'ServicePerimeter' that are allowed to be accessed by sources defined in the corresponding 'IngressFrom'.
	//
	// A request matches if it contains
	// a resource in this list. If '*' is specified for resources,
	// then this 'IngressTo' rule will authorize access to all
	// resources inside the perimeter, provided that the request
	// also matches the 'operations' field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/access_context_manager_service_perimeter#resources AccessContextManagerServicePerimeter#resources}
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// A list of IAM roles that represent the set of operations that the sources specified in the corresponding 'IngressFrom' are allowed to perform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/access_context_manager_service_perimeter#roles AccessContextManagerServicePerimeter#roles}
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
}

