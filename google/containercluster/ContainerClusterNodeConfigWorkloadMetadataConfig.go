package containercluster


type ContainerClusterNodeConfigWorkloadMetadataConfig struct {
	// Mode is the configuration for how to expose metadata to workloads running on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.63.0/docs/resources/container_cluster#mode ContainerCluster#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

