// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionLogConfig struct {
	// Enabled represents whether logging is enabled or not for a connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/integration_connectors_connection#enabled IntegrationConnectorsConnection#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

