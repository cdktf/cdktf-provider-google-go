// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterControlPlaneEndpointsConfigDnsEndpointConfig struct {
	// Controls whether user traffic is allowed over this endpoint.
	//
	// Note that GCP-managed services may still use the endpoint even if this is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/container_cluster#allow_external_traffic ContainerCluster#allow_external_traffic}
	AllowExternalTraffic interface{} `field:"optional" json:"allowExternalTraffic" yaml:"allowExternalTraffic"`
	// The cluster's DNS endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/container_cluster#endpoint ContainerCluster#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

