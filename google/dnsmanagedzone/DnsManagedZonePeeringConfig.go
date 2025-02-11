// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnsmanagedzone


type DnsManagedZonePeeringConfig struct {
	// target_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/dns_managed_zone#target_network DnsManagedZone#target_network}
	TargetNetwork *DnsManagedZonePeeringConfigTargetNetwork `field:"required" json:"targetNetwork" yaml:"targetNetwork"`
}

