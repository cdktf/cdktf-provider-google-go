// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareenginecluster


type VmwareengineClusterAutoscalingSettingsAutoscalingPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/vmwareengine_cluster#autoscale_policy_id VmwareengineCluster#autoscale_policy_id}.
	AutoscalePolicyId *string `field:"required" json:"autoscalePolicyId" yaml:"autoscalePolicyId"`
	// The canonical identifier of the node type to add or remove.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/vmwareengine_cluster#node_type_id VmwareengineCluster#node_type_id}
	NodeTypeId *string `field:"required" json:"nodeTypeId" yaml:"nodeTypeId"`
	// Number of nodes to add to a cluster during a scale-out operation. Must be divisible by 2 for stretched clusters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/vmwareengine_cluster#scale_out_size VmwareengineCluster#scale_out_size}
	ScaleOutSize *float64 `field:"required" json:"scaleOutSize" yaml:"scaleOutSize"`
	// consumed_memory_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/vmwareengine_cluster#consumed_memory_thresholds VmwareengineCluster#consumed_memory_thresholds}
	ConsumedMemoryThresholds *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds `field:"optional" json:"consumedMemoryThresholds" yaml:"consumedMemoryThresholds"`
	// cpu_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/vmwareengine_cluster#cpu_thresholds VmwareengineCluster#cpu_thresholds}
	CpuThresholds *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds `field:"optional" json:"cpuThresholds" yaml:"cpuThresholds"`
	// storage_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/vmwareengine_cluster#storage_thresholds VmwareengineCluster#storage_thresholds}
	StorageThresholds *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds `field:"optional" json:"storageThresholds" yaml:"storageThresholds"`
}

