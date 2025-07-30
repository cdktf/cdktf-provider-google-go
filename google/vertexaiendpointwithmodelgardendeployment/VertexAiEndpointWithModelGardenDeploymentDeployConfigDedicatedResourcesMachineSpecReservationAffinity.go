// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpointwithmodelgardendeployment


type VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecReservationAffinity struct {
	// Specifies the reservation affinity type. Possible values: TYPE_UNSPECIFIED NO_RESERVATION ANY_RESERVATION SPECIFIC_RESERVATION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#reservation_affinity_type VertexAiEndpointWithModelGardenDeployment#reservation_affinity_type}
	ReservationAffinityType *string `field:"required" json:"reservationAffinityType" yaml:"reservationAffinityType"`
	// Corresponds to the label key of a reservation resource.
	//
	// To target a
	// SPECIFIC_RESERVATION by name, use 'compute.googleapis.com/reservation-name'
	// as the key and specify the name of your reservation as its value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#key VertexAiEndpointWithModelGardenDeployment#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Corresponds to the label values of a reservation resource.
	//
	// This must be the
	// full resource name of the reservation or reservation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#values VertexAiEndpointWithModelGardenDeployment#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

