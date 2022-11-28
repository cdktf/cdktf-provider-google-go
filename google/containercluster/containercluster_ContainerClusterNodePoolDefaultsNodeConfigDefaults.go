package containercluster


type ContainerClusterNodePoolDefaultsNodeConfigDefaults struct {
	// Type of logging agent that is used as the default value for node pools in the cluster.
	//
	// Valid values include DEFAULT and MAX_THROUGHPUT.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#logging_variant ContainerCluster#logging_variant}
	LoggingVariant *string `field:"optional" json:"loggingVariant" yaml:"loggingVariant"`
}

