// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainercluster


type EdgecontainerClusterNetworking struct {
	// All pods in the cluster are assigned an RFC1918 IPv4 address from these blocks.
	//
	// Only a single block is supported. This field cannot be changed
	// after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/edgecontainer_cluster#cluster_ipv4_cidr_blocks EdgecontainerCluster#cluster_ipv4_cidr_blocks}
	ClusterIpv4CidrBlocks *[]*string `field:"required" json:"clusterIpv4CidrBlocks" yaml:"clusterIpv4CidrBlocks"`
	// All services in the cluster are assigned an RFC1918 IPv4 address from these blocks.
	//
	// Only a single block is supported. This field cannot be changed
	// after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/edgecontainer_cluster#services_ipv4_cidr_blocks EdgecontainerCluster#services_ipv4_cidr_blocks}
	ServicesIpv4CidrBlocks *[]*string `field:"required" json:"servicesIpv4CidrBlocks" yaml:"servicesIpv4CidrBlocks"`
	// If specified, dual stack mode is enabled and all pods in the cluster are assigned an IPv6 address from these blocks alongside from an IPv4 address.
	//
	// Only a single block is supported. This field cannot be changed
	// after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/edgecontainer_cluster#cluster_ipv6_cidr_blocks EdgecontainerCluster#cluster_ipv6_cidr_blocks}
	ClusterIpv6CidrBlocks *[]*string `field:"optional" json:"clusterIpv6CidrBlocks" yaml:"clusterIpv6CidrBlocks"`
	// If specified, dual stack mode is enabled and all services in the cluster are assigned an IPv6 address from these blocks alongside from an IPv4 address.
	//
	// Only a single block is supported. This field cannot be changed
	// after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/edgecontainer_cluster#services_ipv6_cidr_blocks EdgecontainerCluster#services_ipv6_cidr_blocks}
	ServicesIpv6CidrBlocks *[]*string `field:"optional" json:"servicesIpv6CidrBlocks" yaml:"servicesIpv6CidrBlocks"`
}

