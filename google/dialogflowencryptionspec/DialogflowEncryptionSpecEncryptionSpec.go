// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowencryptionspec


type DialogflowEncryptionSpecEncryptionSpec struct {
	// The name of customer-managed encryption key that is used to secure a resource and its sub-resources.
	//
	// If empty, the resource is secured by the default Google encryption key.
	// Only the key in the same location as this resource is allowed to be used for encryption.
	// Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{key}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dialogflow_encryption_spec#kms_key DialogflowEncryptionSpec#kms_key}
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
}

