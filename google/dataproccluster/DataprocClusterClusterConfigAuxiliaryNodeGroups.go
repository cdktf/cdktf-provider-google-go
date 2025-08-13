// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster


type DataprocClusterClusterConfigAuxiliaryNodeGroups struct {
	// node_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dataproc_cluster#node_group DataprocCluster#node_group}
	NodeGroup interface{} `field:"required" json:"nodeGroup" yaml:"nodeGroup"`
	// A node group ID.
	//
	// Generated if not specified. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of from 3 to 33 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dataproc_cluster#node_group_id DataprocCluster#node_group_id}
	NodeGroupId *string `field:"optional" json:"nodeGroupId" yaml:"nodeGroupId"`
}

