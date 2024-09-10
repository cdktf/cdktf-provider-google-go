// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareengineprivatecloud


type VmwareenginePrivateCloudManagementCluster struct {
	// The user-provided identifier of the new Cluster.
	//
	// The identifier must meet the following requirements:
	//   * Only contains 1-63 alphanumeric characters and hyphens
	//   * Begins with an alphabetical character
	//   * Ends with a non-hyphen character
	//   * Not formatted as a UUID
	//   * Complies with RFC 1034 (https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/vmwareengine_private_cloud#cluster_id VmwareenginePrivateCloud#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// node_type_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/vmwareengine_private_cloud#node_type_configs VmwareenginePrivateCloud#node_type_configs}
	NodeTypeConfigs interface{} `field:"optional" json:"nodeTypeConfigs" yaml:"nodeTypeConfigs"`
	// stretched_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/vmwareengine_private_cloud#stretched_cluster_config VmwareenginePrivateCloud#stretched_cluster_config}
	StretchedClusterConfig *VmwareenginePrivateCloudManagementClusterStretchedClusterConfig `field:"optional" json:"stretchedClusterConfig" yaml:"stretchedClusterConfig"`
}

