// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computereservation


type ComputeReservationSpecificReservation struct {
	// The number of resources that are allocated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/compute_reservation#count ComputeReservation#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// instance_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/compute_reservation#instance_properties ComputeReservation#instance_properties}
	InstanceProperties *ComputeReservationSpecificReservationInstanceProperties `field:"required" json:"instanceProperties" yaml:"instanceProperties"`
}

