package containercluster


type ContainerClusterNetworkPolicy struct {
	// Whether network policy is enabled on the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The selected network policy provider. Defaults to PROVIDER_UNSPECIFIED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#provider ContainerCluster#provider}
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
}

