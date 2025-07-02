// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionConfigVariableEncryptionKeyValue struct {
	// Type of Encryption Key Possible values: ["GOOGLE_MANAGED", "CUSTOMER_MANAGED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/integration_connectors_connection#type IntegrationConnectorsConnection#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The [KMS key name] with which the content of the Operation is encrypted.
	//
	// The
	// expected format: projects/* /locations/* /keyRings/* /cryptoKeys/*.
	// Will be empty string if google managed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/integration_connectors_connection#kms_key_name IntegrationConnectorsConnection#kms_key_name}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
}

