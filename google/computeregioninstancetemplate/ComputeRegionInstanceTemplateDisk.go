// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancetemplate


type ComputeRegionInstanceTemplateDisk struct {
	// The architecture of the image. Allowed values are ARM64 or X86_64.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#architecture ComputeRegionInstanceTemplate#architecture}
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// Whether or not the disk should be auto-deleted. This defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#auto_delete ComputeRegionInstanceTemplate#auto_delete}
	AutoDelete interface{} `field:"optional" json:"autoDelete" yaml:"autoDelete"`
	// Indicates that this is a boot disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#boot ComputeRegionInstanceTemplate#boot}
	Boot interface{} `field:"optional" json:"boot" yaml:"boot"`
	// A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance.
	//
	// If not specified, the server chooses a default device name to apply to this disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#device_name ComputeRegionInstanceTemplate#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// disk_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#disk_encryption_key ComputeRegionInstanceTemplate#disk_encryption_key}
	DiskEncryptionKey *ComputeRegionInstanceTemplateDiskDiskEncryptionKey `field:"optional" json:"diskEncryptionKey" yaml:"diskEncryptionKey"`
	// Name of the disk. When not provided, this defaults to the name of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#disk_name ComputeRegionInstanceTemplate#disk_name}
	DiskName *string `field:"optional" json:"diskName" yaml:"diskName"`
	// The size of the image in gigabytes.
	//
	// If not specified, it will inherit the size of its base image. For SCRATCH disks, the size must be one of 375 or 3000 GB, with a default of 375 GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#disk_size_gb ComputeRegionInstanceTemplate#disk_size_gb}
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// The Google Compute Engine disk type. Such as "pd-ssd", "local-ssd", "pd-balanced" or "pd-standard".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#disk_type ComputeRegionInstanceTemplate#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// A list of features to enable on the guest operating system. Applicable only for bootable images.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#guest_os_features ComputeRegionInstanceTemplate#guest_os_features}
	GuestOsFeatures *[]*string `field:"optional" json:"guestOsFeatures" yaml:"guestOsFeatures"`
	// Specifies the disk interface to use for attaching this disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#interface ComputeRegionInstanceTemplate#interface}
	Interface *string `field:"optional" json:"interface" yaml:"interface"`
	// A set of key/value label pairs to assign to disks,.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#labels ComputeRegionInstanceTemplate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The mode in which to attach this disk, either READ_WRITE or READ_ONLY.
	//
	// If you are attaching or creating a boot disk, this must read-write mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#mode ComputeRegionInstanceTemplate#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Indicates how many IOPS to provision for the disk.
	//
	// This sets the number of I/O operations per second that the disk can handle. For more details, see the [Extreme persistent disk documentation](https://cloud.google.com/compute/docs/disks/extreme-persistent-disk) or the [Hyperdisk documentation](https://cloud.google.com/compute/docs/disks/hyperdisks) depending on the selected disk_type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#provisioned_iops ComputeRegionInstanceTemplate#provisioned_iops}
	ProvisionedIops *float64 `field:"optional" json:"provisionedIops" yaml:"provisionedIops"`
	// Indicates how much throughput to provision for the disk, in MB/s.
	//
	// This sets the amount of data that can be read or written from the disk per second. Values must greater than or equal to 1. For more details, see the [Hyperdisk documentation](https://cloud.google.com/compute/docs/disks/hyperdisks).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#provisioned_throughput ComputeRegionInstanceTemplate#provisioned_throughput}
	ProvisionedThroughput *float64 `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// A map of resource manager tags.
	//
	// Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT & PATCH) when empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#resource_manager_tags ComputeRegionInstanceTemplate#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
	// A list (short name or id) of resource policies to attach to this disk.
	//
	// Currently a max of 1 resource policy is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#resource_policies ComputeRegionInstanceTemplate#resource_policies}
	ResourcePolicies *[]*string `field:"optional" json:"resourcePolicies" yaml:"resourcePolicies"`
	// The name (not self_link) of the disk (such as those managed by google_compute_disk) to attach.
	//
	// ~> Note: Either source or source_image is required when creating a new instance except for when creating a local SSD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#source ComputeRegionInstanceTemplate#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The image from which to initialize this disk.
	//
	// This can be one of: the image's self_link, projects/{project}/global/images/{image}, projects/{project}/global/images/family/{family}, global/images/{image}, global/images/family/{family}, family/{family}, {project}/{family}, {project}/{image}, {family}, or {image}. ~> Note: Either source or source_image is required when creating a new instance except for when creating a local SSD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#source_image ComputeRegionInstanceTemplate#source_image}
	SourceImage *string `field:"optional" json:"sourceImage" yaml:"sourceImage"`
	// source_image_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#source_image_encryption_key ComputeRegionInstanceTemplate#source_image_encryption_key}
	SourceImageEncryptionKey *ComputeRegionInstanceTemplateDiskSourceImageEncryptionKey `field:"optional" json:"sourceImageEncryptionKey" yaml:"sourceImageEncryptionKey"`
	// The source snapshot to create this disk.
	//
	// When creating
	// a new instance, one of initializeParams.sourceSnapshot,
	// initializeParams.sourceImage, or disks.source is
	// required except for local SSD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#source_snapshot ComputeRegionInstanceTemplate#source_snapshot}
	SourceSnapshot *string `field:"optional" json:"sourceSnapshot" yaml:"sourceSnapshot"`
	// source_snapshot_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#source_snapshot_encryption_key ComputeRegionInstanceTemplate#source_snapshot_encryption_key}
	SourceSnapshotEncryptionKey *ComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKey `field:"optional" json:"sourceSnapshotEncryptionKey" yaml:"sourceSnapshotEncryptionKey"`
	// The type of Google Compute Engine disk, can be either "SCRATCH" or "PERSISTENT".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/compute_region_instance_template#type ComputeRegionInstanceTemplate#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

