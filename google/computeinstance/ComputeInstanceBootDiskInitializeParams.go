// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstance


type ComputeInstanceBootDiskInitializeParams struct {
	// The architecture of the disk. One of "X86_64" or "ARM64".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_instance#architecture ComputeInstance#architecture}
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// A flag to enable confidential compute mode on boot disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_instance#enable_confidential_compute ComputeInstance#enable_confidential_compute}
	EnableConfidentialCompute interface{} `field:"optional" json:"enableConfidentialCompute" yaml:"enableConfidentialCompute"`
	// The image from which this disk was initialised.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_instance#image ComputeInstance#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// A set of key/value label pairs assigned to the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_instance#labels ComputeInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Indicates how many IOPS to provision for the disk.
	//
	// This sets the number of I/O operations per second that the disk can handle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_instance#provisioned_iops ComputeInstance#provisioned_iops}
	ProvisionedIops *float64 `field:"optional" json:"provisionedIops" yaml:"provisionedIops"`
	// Indicates how much throughput to provision for the disk.
	//
	// This sets the number of throughput mb per second that the disk can handle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_instance#provisioned_throughput ComputeInstance#provisioned_throughput}
	ProvisionedThroughput *float64 `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// A map of resource manager tags.
	//
	// Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT & PATCH) when empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_instance#resource_manager_tags ComputeInstance#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
	// A list of self_links of resource policies to attach to the instance's boot disk.
	//
	// Modifying this list will cause the instance to recreate. Currently a max of 1 resource policy is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_instance#resource_policies ComputeInstance#resource_policies}
	ResourcePolicies *[]*string `field:"optional" json:"resourcePolicies" yaml:"resourcePolicies"`
	// The size of the image in gigabytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_instance#size ComputeInstance#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The snapshot from which this disk was initialised.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_instance#snapshot ComputeInstance#snapshot}
	Snapshot *string `field:"optional" json:"snapshot" yaml:"snapshot"`
	// source_image_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_instance#source_image_encryption_key ComputeInstance#source_image_encryption_key}
	SourceImageEncryptionKey *ComputeInstanceBootDiskInitializeParamsSourceImageEncryptionKey `field:"optional" json:"sourceImageEncryptionKey" yaml:"sourceImageEncryptionKey"`
	// source_snapshot_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_instance#source_snapshot_encryption_key ComputeInstance#source_snapshot_encryption_key}
	SourceSnapshotEncryptionKey *ComputeInstanceBootDiskInitializeParamsSourceSnapshotEncryptionKey `field:"optional" json:"sourceSnapshotEncryptionKey" yaml:"sourceSnapshotEncryptionKey"`
	// The URL of the storage pool in which the new disk is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_instance#storage_pool ComputeInstance#storage_pool}
	StoragePool *string `field:"optional" json:"storagePool" yaml:"storagePool"`
	// The Google Compute Engine disk type. Such as pd-standard, pd-ssd or pd-balanced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_instance#type ComputeInstance#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

