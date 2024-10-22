// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsmanagedzone


type IntegrationConnectorsManagedZoneTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/integration_connectors_managed_zone#create IntegrationConnectorsManagedZone#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/integration_connectors_managed_zone#delete IntegrationConnectorsManagedZone#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/integration_connectors_managed_zone#update IntegrationConnectorsManagedZone#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

