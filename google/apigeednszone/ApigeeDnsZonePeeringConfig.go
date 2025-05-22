// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeednszone


type ApigeeDnsZonePeeringConfig struct {
	// The name of the producer VPC network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/apigee_dns_zone#target_network_id ApigeeDnsZone#target_network_id}
	TargetNetworkId *string `field:"required" json:"targetNetworkId" yaml:"targetNetworkId"`
	// The ID of the project that contains the producer VPC network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/apigee_dns_zone#target_project_id ApigeeDnsZone#target_project_id}
	TargetProjectId *string `field:"required" json:"targetProjectId" yaml:"targetProjectId"`
}

