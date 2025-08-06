// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterDnsConfig struct {
	// Enable additive VPC scope DNS in a GKE cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/container_cluster#additive_vpc_scope_dns_domain ContainerCluster#additive_vpc_scope_dns_domain}
	AdditiveVpcScopeDnsDomain *string `field:"optional" json:"additiveVpcScopeDnsDomain" yaml:"additiveVpcScopeDnsDomain"`
	// Which in-cluster DNS provider should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/container_cluster#cluster_dns ContainerCluster#cluster_dns}
	ClusterDns *string `field:"optional" json:"clusterDns" yaml:"clusterDns"`
	// The suffix used for all cluster service records.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/container_cluster#cluster_dns_domain ContainerCluster#cluster_dns_domain}
	ClusterDnsDomain *string `field:"optional" json:"clusterDnsDomain" yaml:"clusterDnsDomain"`
	// The scope of access to cluster DNS records.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/container_cluster#cluster_dns_scope ContainerCluster#cluster_dns_scope}
	ClusterDnsScope *string `field:"optional" json:"clusterDnsScope" yaml:"clusterDnsScope"`
}

