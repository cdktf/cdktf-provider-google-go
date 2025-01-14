// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerattachedcluster


type ContainerAttachedClusterProxyConfigKubernetesSecret struct {
	// Name of the kubernetes secret containing the proxy config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/container_attached_cluster#name ContainerAttachedCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace of the kubernetes secret containing the proxy config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/container_attached_cluster#namespace ContainerAttachedCluster#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

