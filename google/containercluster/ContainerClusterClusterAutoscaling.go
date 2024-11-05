// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterClusterAutoscaling struct {
	// auto_provisioning_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/container_cluster#auto_provisioning_defaults ContainerCluster#auto_provisioning_defaults}
	AutoProvisioningDefaults *ContainerClusterClusterAutoscalingAutoProvisioningDefaults `field:"optional" json:"autoProvisioningDefaults" yaml:"autoProvisioningDefaults"`
	// The list of Google Compute Engine zones in which the NodePool's nodes can be created by NAP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/container_cluster#auto_provisioning_locations ContainerCluster#auto_provisioning_locations}
	AutoProvisioningLocations *[]*string `field:"optional" json:"autoProvisioningLocations" yaml:"autoProvisioningLocations"`
	// Configuration options for the Autoscaling profile feature, which lets you choose whether the cluster autoscaler should optimize for resource utilization or resource availability when deciding to remove nodes from a cluster.
	//
	// Can be BALANCED or OPTIMIZE_UTILIZATION. Defaults to BALANCED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/container_cluster#autoscaling_profile ContainerCluster#autoscaling_profile}
	AutoscalingProfile *string `field:"optional" json:"autoscalingProfile" yaml:"autoscalingProfile"`
	// Whether node auto-provisioning is enabled. Resource limits for cpu and memory must be defined to enable node auto-provisioning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// resource_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/container_cluster#resource_limits ContainerCluster#resource_limits}
	ResourceLimits interface{} `field:"optional" json:"resourceLimits" yaml:"resourceLimits"`
}

