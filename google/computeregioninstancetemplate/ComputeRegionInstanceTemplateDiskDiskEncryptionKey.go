// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancetemplate


type ComputeRegionInstanceTemplateDiskDiskEncryptionKey struct {
	// The self link of the encryption key that is stored in Google Cloud KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/compute_region_instance_template#kms_key_self_link ComputeRegionInstanceTemplate#kms_key_self_link}
	KmsKeySelfLink *string `field:"required" json:"kmsKeySelfLink" yaml:"kmsKeySelfLink"`
}

