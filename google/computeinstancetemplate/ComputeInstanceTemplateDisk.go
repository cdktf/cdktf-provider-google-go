package computeinstancetemplate


type ComputeInstanceTemplateDisk struct {
	// Whether or not the disk should be auto-deleted. This defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#auto_delete ComputeInstanceTemplate#auto_delete}
	AutoDelete interface{} `field:"optional" json:"autoDelete" yaml:"autoDelete"`
	// Indicates that this is a boot disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#boot ComputeInstanceTemplate#boot}
	Boot interface{} `field:"optional" json:"boot" yaml:"boot"`
	// A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance.
	//
	// If not specified, the server chooses a default device name to apply to this disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#device_name ComputeInstanceTemplate#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// disk_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#disk_encryption_key ComputeInstanceTemplate#disk_encryption_key}
	DiskEncryptionKey *ComputeInstanceTemplateDiskDiskEncryptionKey `field:"optional" json:"diskEncryptionKey" yaml:"diskEncryptionKey"`
	// Name of the disk. When not provided, this defaults to the name of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#disk_name ComputeInstanceTemplate#disk_name}
	DiskName *string `field:"optional" json:"diskName" yaml:"diskName"`
	// The size of the image in gigabytes.
	//
	// If not specified, it will inherit the size of its base image. For SCRATCH disks, the size must be one of 375 or 3000 GB, with a default of 375 GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#disk_size_gb ComputeInstanceTemplate#disk_size_gb}
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// The Google Compute Engine disk type. Such as "pd-ssd", "local-ssd", "pd-balanced" or "pd-standard".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#disk_type ComputeInstanceTemplate#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// Specifies the disk interface to use for attaching this disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#interface ComputeInstanceTemplate#interface}
	Interface *string `field:"optional" json:"interface" yaml:"interface"`
	// A set of key/value label pairs to assign to disks,.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#labels ComputeInstanceTemplate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The mode in which to attach this disk, either READ_WRITE or READ_ONLY.
	//
	// If you are attaching or creating a boot disk, this must read-write mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#mode ComputeInstanceTemplate#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Indicates how many IOPS to provision for the disk.
	//
	// This sets the number of I/O operations per second that the disk can handle. Values must be between 10,000 and 120,000. For more details, see the [Extreme persistent disk documentation](https://cloud.google.com/compute/docs/disks/extreme-persistent-disk).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#provisioned_iops ComputeInstanceTemplate#provisioned_iops}
	ProvisionedIops *float64 `field:"optional" json:"provisionedIops" yaml:"provisionedIops"`
	// A list (short name or id) of resource policies to attach to this disk.
	//
	// Currently a max of 1 resource policy is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#resource_policies ComputeInstanceTemplate#resource_policies}
	ResourcePolicies *[]*string `field:"optional" json:"resourcePolicies" yaml:"resourcePolicies"`
	// The name (not self_link) of the disk (such as those managed by google_compute_disk) to attach.
	//
	// ~> Note: Either source or source_image is required when creating a new instance except for when creating a local SSD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#source ComputeInstanceTemplate#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The image from which to initialize this disk.
	//
	// This can be one of: the image's self_link, projects/{project}/global/images/{image}, projects/{project}/global/images/family/{family}, global/images/{image}, global/images/family/{family}, family/{family}, {project}/{family}, {project}/{image}, {family}, or {image}. ~> Note: Either source or source_image is required when creating a new instance except for when creating a local SSD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#source_image ComputeInstanceTemplate#source_image}
	SourceImage *string `field:"optional" json:"sourceImage" yaml:"sourceImage"`
	// source_image_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#source_image_encryption_key ComputeInstanceTemplate#source_image_encryption_key}
	SourceImageEncryptionKey *ComputeInstanceTemplateDiskSourceImageEncryptionKey `field:"optional" json:"sourceImageEncryptionKey" yaml:"sourceImageEncryptionKey"`
	// The source snapshot to create this disk.
	//
	// When creating
	// a new instance, one of initializeParams.sourceSnapshot,
	// initializeParams.sourceImage, or disks.source is
	// required except for local SSD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#source_snapshot ComputeInstanceTemplate#source_snapshot}
	SourceSnapshot *string `field:"optional" json:"sourceSnapshot" yaml:"sourceSnapshot"`
	// source_snapshot_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#source_snapshot_encryption_key ComputeInstanceTemplate#source_snapshot_encryption_key}
	SourceSnapshotEncryptionKey *ComputeInstanceTemplateDiskSourceSnapshotEncryptionKey `field:"optional" json:"sourceSnapshotEncryptionKey" yaml:"sourceSnapshotEncryptionKey"`
	// The type of Google Compute Engine disk, can be either "SCRATCH" or "PERSISTENT".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#type ComputeInstanceTemplate#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

