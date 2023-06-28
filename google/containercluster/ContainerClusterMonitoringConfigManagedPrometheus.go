package containercluster


type ContainerClusterMonitoringConfigManagedPrometheus struct {
	// Whether or not the managed collection is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.71.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

