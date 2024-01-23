// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionEventingConfigAdditionalVariable struct {
	// Key for the configVariable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/integration_connectors_connection#key IntegrationConnectorsConnection#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Boolean Value of configVariable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/integration_connectors_connection#boolean_value IntegrationConnectorsConnection#boolean_value}
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// encryption_key_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/integration_connectors_connection#encryption_key_value IntegrationConnectorsConnection#encryption_key_value}
	EncryptionKeyValue *IntegrationConnectorsConnectionEventingConfigAdditionalVariableEncryptionKeyValue `field:"optional" json:"encryptionKeyValue" yaml:"encryptionKeyValue"`
	// Integer Value of configVariable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/integration_connectors_connection#integer_value IntegrationConnectorsConnection#integer_value}
	IntegerValue *float64 `field:"optional" json:"integerValue" yaml:"integerValue"`
	// secret_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/integration_connectors_connection#secret_value IntegrationConnectorsConnection#secret_value}
	SecretValue *IntegrationConnectorsConnectionEventingConfigAdditionalVariableSecretValue `field:"optional" json:"secretValue" yaml:"secretValue"`
	// String Value of configVariabley.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/integration_connectors_connection#string_value IntegrationConnectorsConnection#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

