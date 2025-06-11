// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagementvpcflowlogsconfig


type NetworkManagementVpcFlowLogsConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/network_management_vpc_flow_logs_config#create NetworkManagementVpcFlowLogsConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/network_management_vpc_flow_logs_config#delete NetworkManagementVpcFlowLogsConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/network_management_vpc_flow_logs_config#update NetworkManagementVpcFlowLogsConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

