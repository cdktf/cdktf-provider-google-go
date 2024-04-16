// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanageraccesslevel


type AccessContextManagerAccessLevelBasicConditionsVpcNetworkSources struct {
	// vpc_subnetwork block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.25.0/docs/resources/access_context_manager_access_level#vpc_subnetwork AccessContextManagerAccessLevel#vpc_subnetwork}
	VpcSubnetwork *AccessContextManagerAccessLevelBasicConditionsVpcNetworkSourcesVpcSubnetwork `field:"optional" json:"vpcSubnetwork" yaml:"vpcSubnetwork"`
}

