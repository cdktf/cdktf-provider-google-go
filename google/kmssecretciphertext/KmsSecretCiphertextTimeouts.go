// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmssecretciphertext


type KmsSecretCiphertextTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/kms_secret_ciphertext#create KmsSecretCiphertext#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/kms_secret_ciphertext#delete KmsSecretCiphertext#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

