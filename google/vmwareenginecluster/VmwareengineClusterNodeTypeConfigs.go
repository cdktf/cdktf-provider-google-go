// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareenginecluster


type VmwareengineClusterNodeTypeConfigs struct {
	// The number of nodes of this type in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/vmwareengine_cluster#node_count VmwareengineCluster#node_count}
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/vmwareengine_cluster#node_type_id VmwareengineCluster#node_type_id}.
	NodeTypeId *string `field:"required" json:"nodeTypeId" yaml:"nodeTypeId"`
	// Customized number of cores available to each node of the type.
	//
	// This number must always be one of 'nodeType.availableCustomCoreCounts'.
	// If zero is provided max value from 'nodeType.availableCustomCoreCounts' will be used.
	// Once the customer is created then corecount cannot be changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/vmwareengine_cluster#custom_core_count VmwareengineCluster#custom_core_count}
	CustomCoreCount *float64 `field:"optional" json:"customCoreCount" yaml:"customCoreCount"`
}

