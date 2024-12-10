// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanageraccesslevelcondition


type AccessContextManagerAccessLevelConditionVpcNetworkSources struct {
	// vpc_subnetwork block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/access_context_manager_access_level_condition#vpc_subnetwork AccessContextManagerAccessLevelCondition#vpc_subnetwork}
	VpcSubnetwork *AccessContextManagerAccessLevelConditionVpcNetworkSourcesVpcSubnetwork `field:"optional" json:"vpcSubnetwork" yaml:"vpcSubnetwork"`
}

