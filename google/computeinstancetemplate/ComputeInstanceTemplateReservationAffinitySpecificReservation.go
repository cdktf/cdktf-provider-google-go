// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancetemplate


type ComputeInstanceTemplateReservationAffinitySpecificReservation struct {
	// Corresponds to the label key of a reservation resource.
	//
	// To target a SPECIFIC_RESERVATION by name, specify compute.googleapis.com/reservation-name as the key and specify the name of your reservation as the only value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/compute_instance_template#key ComputeInstanceTemplate#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Corresponds to the label values of a reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/compute_instance_template#values ComputeInstanceTemplate#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

