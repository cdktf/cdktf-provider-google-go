package computedisk


type ComputeDiskAsyncPrimaryDisk struct {
	// Primary disk for asynchronous disk replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_disk#disk ComputeDisk#disk}
	Disk *string `field:"required" json:"disk" yaml:"disk"`
}

