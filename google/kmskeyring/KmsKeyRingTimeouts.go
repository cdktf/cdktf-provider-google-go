// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmskeyring


type KmsKeyRingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/kms_key_ring#create KmsKeyRing#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/kms_key_ring#delete KmsKeyRing#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

