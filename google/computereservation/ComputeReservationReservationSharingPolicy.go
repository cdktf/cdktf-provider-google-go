// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computereservation


type ComputeReservationReservationSharingPolicy struct {
	// Sharing config for all Google Cloud services. Possible values: ["ALLOW_ALL", "DISALLOW_ALL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/compute_reservation#service_share_type ComputeReservation#service_share_type}
	ServiceShareType *string `field:"optional" json:"serviceShareType" yaml:"serviceShareType"`
}

