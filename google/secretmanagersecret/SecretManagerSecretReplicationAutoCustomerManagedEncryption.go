// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretmanagersecret


type SecretManagerSecretReplicationAutoCustomerManagedEncryption struct {
	// The resource name of the Cloud KMS CryptoKey used to encrypt secret payloads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/secret_manager_secret#kms_key_name SecretManagerSecret#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

