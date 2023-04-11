package containernodepool


type ContainerNodePoolNodeConfigEphemeralStorageLocalSsdConfig struct {
	// Number of local SSDs to use to back ephemeral storage.
	//
	// Uses NVMe interfaces. Each local SSD must be 375 or 3000 GB in size, and all local SSDs must share the same size.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_node_pool#local_ssd_count ContainerNodePool#local_ssd_count}
	LocalSsdCount *float64 `field:"required" json:"localSsdCount" yaml:"localSsdCount"`
}

