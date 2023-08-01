package computeregiondisk


type ComputeRegionDiskAsyncPrimaryDisk struct {
	// Primary disk for asynchronous disk replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_disk#disk ComputeRegionDisk#disk}
	Disk *string `field:"required" json:"disk" yaml:"disk"`
}

