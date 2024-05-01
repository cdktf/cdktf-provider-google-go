// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacertificateauthority


type PrivatecaCertificateAuthorityConfigSubjectKeyId struct {
	// The value of the KeyId in lowercase hexidecimal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/privateca_certificate_authority#key_id PrivatecaCertificateAuthority#key_id}
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

