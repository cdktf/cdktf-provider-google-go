package computedisk

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeDiskConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Name of the resource.
	//
	// Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#name ComputeDisk#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// async_primary_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#async_primary_disk ComputeDisk#async_primary_disk}
	AsyncPrimaryDisk *ComputeDiskAsyncPrimaryDisk `field:"optional" json:"asyncPrimaryDisk" yaml:"asyncPrimaryDisk"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#description ComputeDisk#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// disk_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#disk_encryption_key ComputeDisk#disk_encryption_key}
	DiskEncryptionKey *ComputeDiskDiskEncryptionKey `field:"optional" json:"diskEncryptionKey" yaml:"diskEncryptionKey"`
	// guest_os_features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#guest_os_features ComputeDisk#guest_os_features}
	GuestOsFeatures interface{} `field:"optional" json:"guestOsFeatures" yaml:"guestOsFeatures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#id ComputeDisk#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The image from which to initialize this disk.
	//
	// This can be
	// one of: the image's 'self_link', 'projects/{project}/global/images/{image}',
	// 'projects/{project}/global/images/family/{family}', 'global/images/{image}',
	// 'global/images/family/{family}', 'family/{family}', '{project}/{family}',
	// '{project}/{image}', '{family}', or '{image}'. If referred by family, the
	// images names must include the family name. If they don't, use the
	// [google_compute_image data source](/docs/providers/google/d/compute_image.html).
	// For instance, the image 'centos-6-v20180104' includes its family name 'centos-6'.
	// These images can be referred by family name here.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#image ComputeDisk#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Labels to apply to this disk.  A list of key->value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#labels ComputeDisk#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Any applicable license URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#licenses ComputeDisk#licenses}
	Licenses *[]*string `field:"optional" json:"licenses" yaml:"licenses"`
	// Physical block size of the persistent disk, in bytes.
	//
	// If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#physical_block_size_bytes ComputeDisk#physical_block_size_bytes}
	PhysicalBlockSizeBytes *float64 `field:"optional" json:"physicalBlockSizeBytes" yaml:"physicalBlockSizeBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#project ComputeDisk#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Indicates how many IOPS must be provisioned for the disk.
	//
	// Note: Updating currently is only supported by hyperdisk skus without the need to delete and recreate the disk, hyperdisk
	// allows for an update of IOPS every 4 hours. To update your hyperdisk more frequently, you'll need to manually delete and recreate it
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#provisioned_iops ComputeDisk#provisioned_iops}
	ProvisionedIops *float64 `field:"optional" json:"provisionedIops" yaml:"provisionedIops"`
	// Indicates how much Throughput must be provisioned for the disk.
	//
	// Note: Updating currently is only supported by hyperdisk skus without the need to delete and recreate the disk, hyperdisk
	// allows for an update of Throughput every 4 hours. To update your hyperdisk more frequently, you'll need to manually delete and recreate it
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#provisioned_throughput ComputeDisk#provisioned_throughput}
	ProvisionedThroughput *float64 `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// Size of the persistent disk, specified in GB.
	//
	// You can specify this
	// field when creating a persistent disk using the 'image' or
	// 'snapshot' parameter, or specify it alone to create an empty
	// persistent disk.
	//
	// If you specify this field along with 'image' or 'snapshot',
	// the value must not be less than the size of the image
	// or the size of the snapshot.
	//
	// ~>**NOTE** If you change the size, Terraform updates the disk size
	// if upsizing is detected but recreates the disk if downsizing is requested.
	// You can add 'lifecycle.prevent_destroy' in the config to prevent destroying
	// and recreating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#size ComputeDisk#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The source snapshot used to create this disk.
	//
	// You can provide this as
	// a partial or full URL to the resource. If the snapshot is in another
	// project than this disk, you must supply a full URL. For example, the
	// following are valid values:
	//
	// 'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot'
	// 'projects/project/global/snapshots/snapshot'
	// 'global/snapshots/snapshot'
	// 'snapshot'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#snapshot ComputeDisk#snapshot}
	Snapshot *string `field:"optional" json:"snapshot" yaml:"snapshot"`
	// The source disk used to create this disk.
	//
	// You can provide this as a partial or full URL to the resource.
	// For example, the following are valid values:
	//
	// https://www.googleapis.com/compute/v1/projects/{project}/zones/{zone}/disks/{disk}
	// https://www.googleapis.com/compute/v1/projects/{project}/regions/{region}/disks/{disk}
	// projects/{project}/zones/{zone}/disks/{disk}
	// projects/{project}/regions/{region}/disks/{disk}
	// zones/{zone}/disks/{disk}
	// regions/{region}/disks/{disk}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#source_disk ComputeDisk#source_disk}
	SourceDisk *string `field:"optional" json:"sourceDisk" yaml:"sourceDisk"`
	// source_image_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#source_image_encryption_key ComputeDisk#source_image_encryption_key}
	SourceImageEncryptionKey *ComputeDiskSourceImageEncryptionKey `field:"optional" json:"sourceImageEncryptionKey" yaml:"sourceImageEncryptionKey"`
	// source_snapshot_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#source_snapshot_encryption_key ComputeDisk#source_snapshot_encryption_key}
	SourceSnapshotEncryptionKey *ComputeDiskSourceSnapshotEncryptionKey `field:"optional" json:"sourceSnapshotEncryptionKey" yaml:"sourceSnapshotEncryptionKey"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#timeouts ComputeDisk#timeouts}
	Timeouts *ComputeDiskTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// URL of the disk type resource describing which disk type to use to create the disk.
	//
	// Provide this when creating the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#type ComputeDisk#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// A reference to the zone where the disk resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#zone ComputeDisk#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

