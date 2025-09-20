// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregiondisk


type ComputeRegionDiskDiskEncryptionKey struct {
	// The name of the encryption key that is stored in Google Cloud KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_disk#kms_key_name ComputeRegionDisk#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_disk#raw_key ComputeRegionDisk#raw_key}
	RawKey *string `field:"optional" json:"rawKey" yaml:"rawKey"`
	// Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit customer-supplied encryption key to either encrypt or decrypt this resource.
	//
	// You can provide either the rawKey or the rsaEncryptedKey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_disk#rsa_encrypted_key ComputeRegionDisk#rsa_encrypted_key}
	RsaEncryptedKey *string `field:"optional" json:"rsaEncryptedKey" yaml:"rsaEncryptedKey"`
}

