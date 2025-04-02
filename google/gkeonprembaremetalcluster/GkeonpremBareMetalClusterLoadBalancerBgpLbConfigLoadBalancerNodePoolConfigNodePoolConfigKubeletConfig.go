// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigKubeletConfig struct {
	// The maximum size of bursty pulls, temporarily allows pulls to burst to this number, while still not exceeding registry_pull_qps.
	//
	// The value must not be a negative number.
	// Updating this field may impact scalability by changing the amount of
	// traffic produced by image pulls.
	// Defaults to 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/gkeonprem_bare_metal_cluster#registry_burst GkeonpremBareMetalCluster#registry_burst}
	RegistryBurst *float64 `field:"optional" json:"registryBurst" yaml:"registryBurst"`
	// The limit of registry pulls per second.
	//
	// Setting this value to 0 means no limit.
	// Updating this field may impact scalability by changing the amount of
	// traffic produced by image pulls.
	// Defaults to 5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/gkeonprem_bare_metal_cluster#registry_pull_qps GkeonpremBareMetalCluster#registry_pull_qps}
	RegistryPullQps *float64 `field:"optional" json:"registryPullQps" yaml:"registryPullQps"`
	// Prevents the Kubelet from pulling multiple images at a time.
	//
	// We recommend *not* changing the default value on nodes that run docker
	// daemon with version  < 1.9 or an Another Union File System (Aufs) storage
	// backend. Issue https://github.com/kubernetes/kubernetes/issues/10959 has
	// more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/gkeonprem_bare_metal_cluster#serialize_image_pulls_disabled GkeonpremBareMetalCluster#serialize_image_pulls_disabled}
	SerializeImagePullsDisabled interface{} `field:"optional" json:"serializeImagePullsDisabled" yaml:"serializeImagePullsDisabled"`
}

