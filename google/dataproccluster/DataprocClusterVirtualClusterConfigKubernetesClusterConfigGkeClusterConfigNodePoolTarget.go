// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster


type DataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget struct {
	// The target GKE node pool. Format: 'projects/{project}/locations/{location}/clusters/{cluster}/nodePools/{nodePool}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/dataproc_cluster#node_pool DataprocCluster#node_pool}
	NodePool *string `field:"required" json:"nodePool" yaml:"nodePool"`
	// The roles associated with the GKE node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/dataproc_cluster#roles DataprocCluster#roles}
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// node_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/dataproc_cluster#node_pool_config DataprocCluster#node_pool_config}
	NodePoolConfig *DataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig `field:"optional" json:"nodePoolConfig" yaml:"nodePoolConfig"`
}

