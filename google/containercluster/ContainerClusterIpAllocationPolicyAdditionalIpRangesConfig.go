// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterIpAllocationPolicyAdditionalIpRangesConfig struct {
	// Name of the subnetwork. This can be the full path of the subnetwork or just the name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/container_cluster#subnetwork ContainerCluster#subnetwork}
	Subnetwork *string `field:"required" json:"subnetwork" yaml:"subnetwork"`
	// List of secondary ranges names within this subnetwork that can be used for pod IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/container_cluster#pod_ipv4_range_names ContainerCluster#pod_ipv4_range_names}
	PodIpv4RangeNames *[]*string `field:"optional" json:"podIpv4RangeNames" yaml:"podIpv4RangeNames"`
}

