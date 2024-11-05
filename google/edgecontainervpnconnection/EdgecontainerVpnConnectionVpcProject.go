// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainervpnconnection


type EdgecontainerVpnConnectionVpcProject struct {
	// The project of the VPC to connect to. If not specified, it is the same as the cluster project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/edgecontainer_vpn_connection#project_id EdgecontainerVpnConnection#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

