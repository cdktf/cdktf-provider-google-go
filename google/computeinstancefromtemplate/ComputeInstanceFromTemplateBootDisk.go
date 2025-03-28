// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancefromtemplate


type ComputeInstanceFromTemplateBootDisk struct {
	// Whether the disk will be auto-deleted when the instance is deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/compute_instance_from_template#auto_delete ComputeInstanceFromTemplate#auto_delete}
	AutoDelete interface{} `field:"optional" json:"autoDelete" yaml:"autoDelete"`
	// Name with which attached disk will be accessible under /dev/disk/by-id/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/compute_instance_from_template#device_name ComputeInstanceFromTemplate#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// A 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to encrypt this disk.
	//
	// Only one of kms_key_self_link and disk_encryption_key_raw may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/compute_instance_from_template#disk_encryption_key_raw ComputeInstanceFromTemplate#disk_encryption_key_raw}
	DiskEncryptionKeyRaw *string `field:"optional" json:"diskEncryptionKeyRaw" yaml:"diskEncryptionKeyRaw"`
	// A list of features to enable on the guest operating system. Applicable only for bootable images.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/compute_instance_from_template#guest_os_features ComputeInstanceFromTemplate#guest_os_features}
	GuestOsFeatures *[]*string `field:"optional" json:"guestOsFeatures" yaml:"guestOsFeatures"`
	// initialize_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/compute_instance_from_template#initialize_params ComputeInstanceFromTemplate#initialize_params}
	InitializeParams *ComputeInstanceFromTemplateBootDiskInitializeParams `field:"optional" json:"initializeParams" yaml:"initializeParams"`
	// The disk interface used for attaching this disk.
	//
	// One of SCSI or NVME. (This field is shared with attached_disk and only used for specific cases, please don't specify this field without advice from Google.)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/compute_instance_from_template#interface ComputeInstanceFromTemplate#interface}
	Interface *string `field:"optional" json:"interface" yaml:"interface"`
	// The self_link of the encryption key that is stored in Google Cloud KMS to encrypt this disk.
	//
	// Only one of kms_key_self_link and disk_encryption_key_raw may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/compute_instance_from_template#kms_key_self_link ComputeInstanceFromTemplate#kms_key_self_link}
	KmsKeySelfLink *string `field:"optional" json:"kmsKeySelfLink" yaml:"kmsKeySelfLink"`
	// Read/write mode for the disk. One of "READ_ONLY" or "READ_WRITE".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/compute_instance_from_template#mode ComputeInstanceFromTemplate#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The name or self_link of the disk attached to this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/compute_instance_from_template#source ComputeInstanceFromTemplate#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}

