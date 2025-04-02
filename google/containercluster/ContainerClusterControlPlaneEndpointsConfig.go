// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterControlPlaneEndpointsConfig struct {
	// dns_endpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/container_cluster#dns_endpoint_config ContainerCluster#dns_endpoint_config}
	DnsEndpointConfig *ContainerClusterControlPlaneEndpointsConfigDnsEndpointConfig `field:"optional" json:"dnsEndpointConfig" yaml:"dnsEndpointConfig"`
	// ip_endpoints_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/container_cluster#ip_endpoints_config ContainerCluster#ip_endpoints_config}
	IpEndpointsConfig *ContainerClusterControlPlaneEndpointsConfigIpEndpointsConfig `field:"optional" json:"ipEndpointsConfig" yaml:"ipEndpointsConfig"`
}

