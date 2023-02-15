package containercluster


type ContainerClusterPrivateClusterConfigMasterGlobalAccessConfig struct {
	// Whether the cluster master is accessible globally or not.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

