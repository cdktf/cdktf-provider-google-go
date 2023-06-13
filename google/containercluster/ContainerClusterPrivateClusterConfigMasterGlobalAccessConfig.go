package containercluster


type ContainerClusterPrivateClusterConfigMasterGlobalAccessConfig struct {
	// Whether the cluster master is accessible globally or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

