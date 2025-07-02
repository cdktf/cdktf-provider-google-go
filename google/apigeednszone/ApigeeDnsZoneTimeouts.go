// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeednszone


type ApigeeDnsZoneTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/apigee_dns_zone#create ApigeeDnsZone#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/apigee_dns_zone#delete ApigeeDnsZone#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

