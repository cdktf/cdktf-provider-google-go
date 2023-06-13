package computeregiondisk

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionDiskConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_disk#name ComputeRegionDisk#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// URLs of the zones where the disk should be replicated to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_disk#replica_zones ComputeRegionDisk#replica_zones}
	ReplicaZones *[]*string `field:"required" json:"replicaZones" yaml:"replicaZones"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_disk#description ComputeRegionDisk#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// disk_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_disk#disk_encryption_key ComputeRegionDisk#disk_encryption_key}
	DiskEncryptionKey *ComputeRegionDiskDiskEncryptionKey `field:"optional" json:"diskEncryptionKey" yaml:"diskEncryptionKey"`
	// guest_os_features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_disk#guest_os_features ComputeRegionDisk#guest_os_features}
	GuestOsFeatures interface{} `field:"optional" json:"guestOsFeatures" yaml:"guestOsFeatures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_disk#id ComputeRegionDisk#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels to apply to this disk.  A list of key->value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_disk#labels ComputeRegionDisk#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Any applicable license URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_disk#licenses ComputeRegionDisk#licenses}
	Licenses *[]*string `field:"optional" json:"licenses" yaml:"licenses"`
	// Physical block size of the persistent disk, in bytes.
	//
	// If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_disk#physical_block_size_bytes ComputeRegionDisk#physical_block_size_bytes}
	PhysicalBlockSizeBytes *float64 `field:"optional" json:"physicalBlockSizeBytes" yaml:"physicalBlockSizeBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_disk#project ComputeRegionDisk#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A reference to the region where the disk resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_disk#region ComputeRegionDisk#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Size of the persistent disk, specified in GB.
	//
	// You can specify this
	// field when creating a persistent disk using the sourceImage or
	// sourceSnapshot parameter, or specify it alone to create an empty
	// persistent disk.
	//
	// If you specify this field along with sourceImage or sourceSnapshot,
	// the value of sizeGb must not be less than the size of the sourceImage
	// or the size of the snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_disk#size ComputeRegionDisk#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The source snapshot used to create this disk.
	//
	// You can provide this as
	// a partial or full URL to the resource. For example, the following are
	// valid values:
	//
	// 'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot'
	// 'projects/project/global/snapshots/snapshot'
	// 'global/snapshots/snapshot'
	// 'snapshot'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_disk#snapshot ComputeRegionDisk#snapshot}
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_disk#source_disk ComputeRegionDisk#source_disk}
	SourceDisk *string `field:"optional" json:"sourceDisk" yaml:"sourceDisk"`
	// source_snapshot_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_disk#source_snapshot_encryption_key ComputeRegionDisk#source_snapshot_encryption_key}
	SourceSnapshotEncryptionKey *ComputeRegionDiskSourceSnapshotEncryptionKey `field:"optional" json:"sourceSnapshotEncryptionKey" yaml:"sourceSnapshotEncryptionKey"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_disk#timeouts ComputeRegionDisk#timeouts}
	Timeouts *ComputeRegionDiskTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// URL of the disk type resource describing which disk type to use to create the disk.
	//
	// Provide this when creating the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_disk#type ComputeRegionDisk#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

