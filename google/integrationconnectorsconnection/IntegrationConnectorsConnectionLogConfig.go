// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionLogConfig struct {
	// Enabled represents whether logging is enabled or not for a connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/integration_connectors_connection#enabled IntegrationConnectorsConnection#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Log configuration level. Possible values: ["LOG_LEVEL_UNSPECIFIED", "ERROR", "INFO", "DEBUG"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/integration_connectors_connection#level IntegrationConnectorsConnection#level}
	Level *string `field:"optional" json:"level" yaml:"level"`
}

