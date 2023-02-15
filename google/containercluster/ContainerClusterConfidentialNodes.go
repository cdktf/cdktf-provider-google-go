package containercluster


type ContainerClusterConfidentialNodes struct {
	// Whether Confidential Nodes feature is enabled for all nodes in this cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

