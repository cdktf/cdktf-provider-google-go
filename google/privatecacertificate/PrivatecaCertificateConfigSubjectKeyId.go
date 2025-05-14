// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacertificate


type PrivatecaCertificateConfigSubjectKeyId struct {
	// The value of the KeyId in lowercase hexadecimal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/privateca_certificate#key_id PrivatecaCertificate#key_id}
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

