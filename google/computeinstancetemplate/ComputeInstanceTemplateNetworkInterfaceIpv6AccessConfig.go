// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancetemplate


type ComputeInstanceTemplateNetworkInterfaceIpv6AccessConfig struct {
	// The service-level to be provided for IPv6 traffic when the subnet has an external subnet.
	//
	// Only PREMIUM tier is valid for IPv6
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/compute_instance_template#network_tier ComputeInstanceTemplate#network_tier}
	NetworkTier *string `field:"required" json:"networkTier" yaml:"networkTier"`
}

