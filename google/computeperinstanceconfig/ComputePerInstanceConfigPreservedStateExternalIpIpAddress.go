// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeperinstanceconfig


type ComputePerInstanceConfigPreservedStateExternalIpIpAddress struct {
	// The URL of the reservation for this IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/compute_per_instance_config#address ComputePerInstanceConfig#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
}

