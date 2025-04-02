// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareengineprivatecloud


type VmwareenginePrivateCloudManagementClusterAutoscalingSettings struct {
	// autoscaling_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/vmwareengine_private_cloud#autoscaling_policies VmwareenginePrivateCloud#autoscaling_policies}
	AutoscalingPolicies interface{} `field:"required" json:"autoscalingPolicies" yaml:"autoscalingPolicies"`
	// The minimum duration between consecutive autoscale operations.
	//
	// It starts once addition or removal of nodes is fully completed.
	// Minimum cool down period is 30m.
	// Cool down period must be in whole minutes (for example, 30m, 31m, 50m).
	// Mandatory for successful addition of autoscaling settings in cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/vmwareengine_private_cloud#cool_down_period VmwareenginePrivateCloud#cool_down_period}
	CoolDownPeriod *string `field:"optional" json:"coolDownPeriod" yaml:"coolDownPeriod"`
	// Maximum number of nodes of any type in a cluster. Mandatory for successful addition of autoscaling settings in cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/vmwareengine_private_cloud#max_cluster_node_count VmwareenginePrivateCloud#max_cluster_node_count}
	MaxClusterNodeCount *float64 `field:"optional" json:"maxClusterNodeCount" yaml:"maxClusterNodeCount"`
	// Minimum number of nodes of any type in a cluster. Mandatory for successful addition of autoscaling settings in cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/vmwareengine_private_cloud#min_cluster_node_count VmwareenginePrivateCloud#min_cluster_node_count}
	MinClusterNodeCount *float64 `field:"optional" json:"minClusterNodeCount" yaml:"minClusterNodeCount"`
}

