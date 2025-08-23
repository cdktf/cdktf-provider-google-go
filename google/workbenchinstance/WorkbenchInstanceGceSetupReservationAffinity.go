// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workbenchinstance


type WorkbenchInstanceGceSetupReservationAffinity struct {
	// Specifies the type of reservation from which this instance can consume resources: RESERVATION_ANY (default), RESERVATION_SPECIFIC, or RESERVATION_NONE.
	//
	// Possible values: ["RESERVATION_NONE", "RESERVATION_ANY", "RESERVATION_SPECIFIC"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/workbench_instance#consume_reservation_type WorkbenchInstance#consume_reservation_type}
	ConsumeReservationType *string `field:"optional" json:"consumeReservationType" yaml:"consumeReservationType"`
	// Corresponds to the label key of a reservation resource.
	//
	// To target a
	// RESERVATION_SPECIFIC by name, use compute.googleapis.com/reservation-name
	// as the key and specify the name of your reservation as its value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/workbench_instance#key WorkbenchInstance#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Corresponds to the label values of a reservation resource.
	//
	// This can be
	// either a name to a reservation in the same project or
	// "projects/different-project/reservations/some-reservation-name"
	// to target a shared reservation in the same zone but in a different project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/workbench_instance#values WorkbenchInstance#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

