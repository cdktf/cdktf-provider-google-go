// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionAuthConfigUserPassword struct {
	// Username for Authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/integration_connectors_connection#username IntegrationConnectorsConnection#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/integration_connectors_connection#password IntegrationConnectorsConnection#password}
	Password *IntegrationConnectorsConnectionAuthConfigUserPasswordPassword `field:"optional" json:"password" yaml:"password"`
}

