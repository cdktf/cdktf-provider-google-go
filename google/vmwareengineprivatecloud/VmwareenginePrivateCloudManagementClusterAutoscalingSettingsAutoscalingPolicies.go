// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareengineprivatecloud


type VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/vmwareengine_private_cloud#autoscale_policy_id VmwareenginePrivateCloud#autoscale_policy_id}.
	AutoscalePolicyId *string `field:"required" json:"autoscalePolicyId" yaml:"autoscalePolicyId"`
	// The canonical identifier of the node type to add or remove.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/vmwareengine_private_cloud#node_type_id VmwareenginePrivateCloud#node_type_id}
	NodeTypeId *string `field:"required" json:"nodeTypeId" yaml:"nodeTypeId"`
	// Number of nodes to add to a cluster during a scale-out operation. Must be divisible by 2 for stretched clusters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/vmwareengine_private_cloud#scale_out_size VmwareenginePrivateCloud#scale_out_size}
	ScaleOutSize *float64 `field:"required" json:"scaleOutSize" yaml:"scaleOutSize"`
	// consumed_memory_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/vmwareengine_private_cloud#consumed_memory_thresholds VmwareenginePrivateCloud#consumed_memory_thresholds}
	ConsumedMemoryThresholds *VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds `field:"optional" json:"consumedMemoryThresholds" yaml:"consumedMemoryThresholds"`
	// cpu_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/vmwareengine_private_cloud#cpu_thresholds VmwareenginePrivateCloud#cpu_thresholds}
	CpuThresholds *VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds `field:"optional" json:"cpuThresholds" yaml:"cpuThresholds"`
	// storage_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/vmwareengine_private_cloud#storage_thresholds VmwareenginePrivateCloud#storage_thresholds}
	StorageThresholds *VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds `field:"optional" json:"storageThresholds" yaml:"storageThresholds"`
}

