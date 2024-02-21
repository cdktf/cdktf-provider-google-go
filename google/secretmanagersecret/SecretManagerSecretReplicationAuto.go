// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretmanagersecret


type SecretManagerSecretReplicationAuto struct {
	// customer_managed_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/secret_manager_secret#customer_managed_encryption SecretManagerSecret#customer_managed_encryption}
	CustomerManagedEncryption *SecretManagerSecretReplicationAutoCustomerManagedEncryption `field:"optional" json:"customerManagedEncryption" yaml:"customerManagedEncryption"`
}

