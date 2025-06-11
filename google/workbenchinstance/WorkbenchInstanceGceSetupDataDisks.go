// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workbenchinstance


type WorkbenchInstanceGceSetupDataDisks struct {
	// Optional. Input only. Disk encryption method used on the boot and data disks, defaults to GMEK. Possible values: ["GMEK", "CMEK"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/workbench_instance#disk_encryption WorkbenchInstance#disk_encryption}
	DiskEncryption *string `field:"optional" json:"diskEncryption" yaml:"diskEncryption"`
	// Optional.
	//
	// The size of the disk in GB attached to this VM instance,
	// up to a maximum of 64000 GB (64 TB). If not specified, this defaults to
	// 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/workbench_instance#disk_size_gb WorkbenchInstance#disk_size_gb}
	DiskSizeGb *string `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Optional. Input only. Indicates the type of the disk. Possible values: ["PD_STANDARD", "PD_SSD", "PD_BALANCED", "PD_EXTREME"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/workbench_instance#disk_type WorkbenchInstance#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// 'Optional.
	//
	// The KMS key used to encrypt the disks,
	// only applicable if disk_encryption is CMEK. Format: 'projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}'
	// Learn more about using your own encryption keys.'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/workbench_instance#kms_key WorkbenchInstance#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

