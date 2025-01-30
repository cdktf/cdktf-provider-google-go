// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterControlPlaneEndpointsConfig struct {
	// dns_endpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/container_cluster#dns_endpoint_config ContainerCluster#dns_endpoint_config}
	DnsEndpointConfig *ContainerClusterControlPlaneEndpointsConfigDnsEndpointConfig `field:"optional" json:"dnsEndpointConfig" yaml:"dnsEndpointConfig"`
}

