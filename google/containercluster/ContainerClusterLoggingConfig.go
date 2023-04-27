package containercluster


type ContainerClusterLoggingConfig struct {
	// GKE components exposing logs. Valid values include SYSTEM_COMPONENTS, APISERVER, CONTROLLER_MANAGER, SCHEDULER, and WORKLOADS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.63.1/docs/resources/container_cluster#enable_components ContainerCluster#enable_components}
	EnableComponents *[]*string `field:"required" json:"enableComponents" yaml:"enableComponents"`
}

