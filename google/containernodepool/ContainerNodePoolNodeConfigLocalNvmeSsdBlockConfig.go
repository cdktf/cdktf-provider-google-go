package containernodepool


type ContainerNodePoolNodeConfigLocalNvmeSsdBlockConfig struct {
	// Number of raw-block local NVMe SSD disks to be attached to the node.
	//
	// Each local SSD is 375 GB in size.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_node_pool#local_ssd_count ContainerNodePool#local_ssd_count}
	LocalSsdCount *float64 `field:"required" json:"localSsdCount" yaml:"localSsdCount"`
}

