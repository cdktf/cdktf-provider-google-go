// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computereservation


type ComputeReservationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.1.0/docs/resources/compute_reservation#create ComputeReservation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.1.0/docs/resources/compute_reservation#delete ComputeReservation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.1.0/docs/resources/compute_reservation#update ComputeReservation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

