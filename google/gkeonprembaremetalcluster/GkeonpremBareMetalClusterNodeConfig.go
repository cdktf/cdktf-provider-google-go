// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterNodeConfig struct {
	// The available runtimes that can be used to run containers in a Bare Metal User Cluster.
	//
	// Possible values: ["CONTAINER_RUNTIME_UNSPECIFIED", "DOCKER", "CONTAINERD"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/gkeonprem_bare_metal_cluster#container_runtime GkeonpremBareMetalCluster#container_runtime}
	ContainerRuntime *string `field:"optional" json:"containerRuntime" yaml:"containerRuntime"`
	// The maximum number of pods a node can run.
	//
	// The size of the CIDR range
	// assigned to the node will be derived from this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/gkeonprem_bare_metal_cluster#max_pods_per_node GkeonpremBareMetalCluster#max_pods_per_node}
	MaxPodsPerNode *float64 `field:"optional" json:"maxPodsPerNode" yaml:"maxPodsPerNode"`
}

