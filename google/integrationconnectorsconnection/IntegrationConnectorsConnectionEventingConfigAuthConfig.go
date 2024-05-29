// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionEventingConfigAuthConfig struct {
	// authType of the Connection Possible values: ["USER_PASSWORD"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/integration_connectors_connection#auth_type IntegrationConnectorsConnection#auth_type}
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// user_password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/integration_connectors_connection#user_password IntegrationConnectorsConnection#user_password}
	UserPassword *IntegrationConnectorsConnectionEventingConfigAuthConfigUserPassword `field:"required" json:"userPassword" yaml:"userPassword"`
	// additional_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/integration_connectors_connection#additional_variable IntegrationConnectorsConnection#additional_variable}
	AdditionalVariable interface{} `field:"optional" json:"additionalVariable" yaml:"additionalVariable"`
	// The type of authentication configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/integration_connectors_connection#auth_key IntegrationConnectorsConnection#auth_key}
	AuthKey *string `field:"optional" json:"authKey" yaml:"authKey"`
}

