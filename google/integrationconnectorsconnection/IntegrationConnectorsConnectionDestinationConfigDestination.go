// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionDestinationConfigDestination struct {
	// For publicly routable host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/integration_connectors_connection#host IntegrationConnectorsConnection#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The port is the target port number that is accepted by the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/integration_connectors_connection#port IntegrationConnectorsConnection#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// PSC service attachments. Format: projects/* /regions/* /serviceAttachments/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/integration_connectors_connection#service_attachment IntegrationConnectorsConnection#service_attachment}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	ServiceAttachment *string `field:"optional" json:"serviceAttachment" yaml:"serviceAttachment"`
}

