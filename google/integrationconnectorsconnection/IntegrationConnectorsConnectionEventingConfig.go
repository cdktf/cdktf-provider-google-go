// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionEventingConfig struct {
	// registration_destination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/integration_connectors_connection#registration_destination_config IntegrationConnectorsConnection#registration_destination_config}
	RegistrationDestinationConfig *IntegrationConnectorsConnectionEventingConfigRegistrationDestinationConfig `field:"required" json:"registrationDestinationConfig" yaml:"registrationDestinationConfig"`
	// additional_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/integration_connectors_connection#additional_variable IntegrationConnectorsConnection#additional_variable}
	AdditionalVariable interface{} `field:"optional" json:"additionalVariable" yaml:"additionalVariable"`
	// auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/integration_connectors_connection#auth_config IntegrationConnectorsConnection#auth_config}
	AuthConfig *IntegrationConnectorsConnectionEventingConfigAuthConfig `field:"optional" json:"authConfig" yaml:"authConfig"`
	// Enrichment Enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/integration_connectors_connection#enrichment_enabled IntegrationConnectorsConnection#enrichment_enabled}
	EnrichmentEnabled interface{} `field:"optional" json:"enrichmentEnabled" yaml:"enrichmentEnabled"`
}

