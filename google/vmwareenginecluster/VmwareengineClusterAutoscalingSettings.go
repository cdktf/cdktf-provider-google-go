// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareenginecluster


type VmwareengineClusterAutoscalingSettings struct {
	// autoscaling_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/vmwareengine_cluster#autoscaling_policies VmwareengineCluster#autoscaling_policies}
	AutoscalingPolicies interface{} `field:"required" json:"autoscalingPolicies" yaml:"autoscalingPolicies"`
	// The minimum duration between consecutive autoscale operations.
	//
	// It starts once addition or removal of nodes is fully completed.
	// Minimum cool down period is 30m.
	// Cool down period must be in whole minutes (for example, 30m, 31m, 50m).
	// Mandatory for successful addition of autoscaling settings in cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/vmwareengine_cluster#cool_down_period VmwareengineCluster#cool_down_period}
	CoolDownPeriod *string `field:"optional" json:"coolDownPeriod" yaml:"coolDownPeriod"`
	// Maximum number of nodes of any type in a cluster. Mandatory for successful addition of autoscaling settings in cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/vmwareengine_cluster#max_cluster_node_count VmwareengineCluster#max_cluster_node_count}
	MaxClusterNodeCount *float64 `field:"optional" json:"maxClusterNodeCount" yaml:"maxClusterNodeCount"`
	// Minimum number of nodes of any type in a cluster. Mandatory for successful addition of autoscaling settings in cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/vmwareengine_cluster#min_cluster_node_count VmwareengineCluster#min_cluster_node_count}
	MinClusterNodeCount *float64 `field:"optional" json:"minClusterNodeCount" yaml:"minClusterNodeCount"`
}

