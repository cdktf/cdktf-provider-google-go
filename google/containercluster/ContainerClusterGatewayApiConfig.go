// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterGatewayApiConfig struct {
	// The Gateway API release channel to use for Gateway API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/container_cluster#channel ContainerCluster#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
}

