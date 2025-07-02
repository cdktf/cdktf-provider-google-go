// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterClusterAutoscalingResourceLimits struct {
	// Maximum amount of the resource in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/container_cluster#maximum ContainerCluster#maximum}
	Maximum *float64 `field:"required" json:"maximum" yaml:"maximum"`
	// The type of the resource.
	//
	// For example, cpu and memory. See the guide to using Node Auto-Provisioning for a list of types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/container_cluster#resource_type ContainerCluster#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Minimum amount of the resource in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/container_cluster#minimum ContainerCluster#minimum}
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
}

