package containerattachedcluster


type ContainerAttachedClusterLoggingConfigComponentConfig struct {
	// The components to be enabled. Possible values: ["SYSTEM_COMPONENTS", "WORKLOADS"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_attached_cluster#enable_components ContainerAttachedCluster#enable_components}
	EnableComponents *[]*string `field:"optional" json:"enableComponents" yaml:"enableComponents"`
}

