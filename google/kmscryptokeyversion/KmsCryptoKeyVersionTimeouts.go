// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmscryptokeyversion


type KmsCryptoKeyVersionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/kms_crypto_key_version#create KmsCryptoKeyVersion#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/kms_crypto_key_version#delete KmsCryptoKeyVersion#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/kms_crypto_key_version#update KmsCryptoKeyVersion#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

