// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionDestinationConfig struct {
	// The key is the destination identifier that is supported by the Connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/integration_connectors_connection#key IntegrationConnectorsConnection#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/integration_connectors_connection#destination IntegrationConnectorsConnection#destination}
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
}

