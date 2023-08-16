package computediskasyncreplication


type ComputeDiskAsyncReplicationSecondaryDisk struct {
	// Secondary disk for asynchronous replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_disk_async_replication#disk ComputeDiskAsyncReplication#disk}
	Disk *string `field:"required" json:"disk" yaml:"disk"`
}

