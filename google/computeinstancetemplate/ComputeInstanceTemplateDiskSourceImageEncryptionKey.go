// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancetemplate


type ComputeInstanceTemplateDiskSourceImageEncryptionKey struct {
	// The self link of the encryption key that is stored in Google Cloud KMS.
	//
	// Only one of kms_key_self_link, rsa_encrypted_key and raw_key may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_instance_template#kms_key_self_link ComputeInstanceTemplate#kms_key_self_link}
	KmsKeySelfLink *string `field:"optional" json:"kmsKeySelfLink" yaml:"kmsKeySelfLink"`
	// The service account being used for the encryption request for the given KMS key.
	//
	// If absent, the Compute
	// Engine default service account is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_instance_template#kms_key_service_account ComputeInstanceTemplate#kms_key_service_account}
	KmsKeyServiceAccount *string `field:"optional" json:"kmsKeyServiceAccount" yaml:"kmsKeyServiceAccount"`
	// Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.
	//
	// Only one of kms_key_self_link, rsa_encrypted_key and raw_key may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_instance_template#raw_key ComputeInstanceTemplate#raw_key}
	RawKey *string `field:"optional" json:"rawKey" yaml:"rawKey"`
	// Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit customer-supplied encryption key to either encrypt or decrypt this resource.
	//
	// Only one of kms_key_self_link, rsa_encrypted_key and raw_key may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_instance_template#rsa_encrypted_key ComputeInstanceTemplate#rsa_encrypted_key}
	RsaEncryptedKey *string `field:"optional" json:"rsaEncryptedKey" yaml:"rsaEncryptedKey"`
}

