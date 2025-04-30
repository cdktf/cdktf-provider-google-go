// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster


type DataprocClusterClusterConfigAuxiliaryNodeGroupsNodeGroup struct {
	// Node group roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/dataproc_cluster#roles DataprocCluster#roles}
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// node_group_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/dataproc_cluster#node_group_config DataprocCluster#node_group_config}
	NodeGroupConfig *DataprocClusterClusterConfigAuxiliaryNodeGroupsNodeGroupNodeGroupConfig `field:"optional" json:"nodeGroupConfig" yaml:"nodeGroupConfig"`
}

