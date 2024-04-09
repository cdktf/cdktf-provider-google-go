// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionEventingConfigRegistrationDestinationConfigDestination struct {
	// Host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.24.0/docs/resources/integration_connectors_connection#host IntegrationConnectorsConnection#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.24.0/docs/resources/integration_connectors_connection#port IntegrationConnectorsConnection#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Service Attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.24.0/docs/resources/integration_connectors_connection#service_attachment IntegrationConnectorsConnection#service_attachment}
	ServiceAttachment *string `field:"optional" json:"serviceAttachment" yaml:"serviceAttachment"`
}

