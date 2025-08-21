// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computereservation


type ComputeReservationDeleteAfterDuration struct {
	// Number of nanoseconds for the auto-delete duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/compute_reservation#nanos ComputeReservation#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
	// Number of seconds for the auto-delete duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/compute_reservation#seconds ComputeReservation#seconds}
	Seconds *string `field:"optional" json:"seconds" yaml:"seconds"`
}

