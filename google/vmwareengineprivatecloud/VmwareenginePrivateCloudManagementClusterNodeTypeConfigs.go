// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareengineprivatecloud


type VmwareenginePrivateCloudManagementClusterNodeTypeConfigs struct {
	// The number of nodes of this type in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/vmwareengine_private_cloud#node_count VmwareenginePrivateCloud#node_count}
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/vmwareengine_private_cloud#node_type_id VmwareenginePrivateCloud#node_type_id}.
	NodeTypeId *string `field:"required" json:"nodeTypeId" yaml:"nodeTypeId"`
	// Customized number of cores available to each node of the type.
	//
	// This number must always be one of 'nodeType.availableCustomCoreCounts'.
	// If zero is provided max value from 'nodeType.availableCustomCoreCounts' will be used.
	// This cannot be changed once the PrivateCloud is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/vmwareengine_private_cloud#custom_core_count VmwareenginePrivateCloud#custom_core_count}
	CustomCoreCount *float64 `field:"optional" json:"customCoreCount" yaml:"customCoreCount"`
}

