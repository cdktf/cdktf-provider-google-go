// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeterdryrunegresspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessContextManagerServicePerimeterDryRunEgressPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the Service Perimeter to add this resource to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/access_context_manager_service_perimeter_dry_run_egress_policy#perimeter AccessContextManagerServicePerimeterDryRunEgressPolicy#perimeter}
	Perimeter *string `field:"required" json:"perimeter" yaml:"perimeter"`
	// egress_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/access_context_manager_service_perimeter_dry_run_egress_policy#egress_from AccessContextManagerServicePerimeterDryRunEgressPolicy#egress_from}
	EgressFrom *AccessContextManagerServicePerimeterDryRunEgressPolicyEgressFrom `field:"optional" json:"egressFrom" yaml:"egressFrom"`
	// egress_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/access_context_manager_service_perimeter_dry_run_egress_policy#egress_to AccessContextManagerServicePerimeterDryRunEgressPolicy#egress_to}
	EgressTo *AccessContextManagerServicePerimeterDryRunEgressPolicyEgressTo `field:"optional" json:"egressTo" yaml:"egressTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/access_context_manager_service_perimeter_dry_run_egress_policy#id AccessContextManagerServicePerimeterDryRunEgressPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/access_context_manager_service_perimeter_dry_run_egress_policy#timeouts AccessContextManagerServicePerimeterDryRunEgressPolicy#timeouts}
	Timeouts *AccessContextManagerServicePerimeterDryRunEgressPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

