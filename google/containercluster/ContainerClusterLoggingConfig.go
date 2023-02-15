package containercluster


type ContainerClusterLoggingConfig struct {
	// GKE components exposing logs. Valid values include SYSTEM_COMPONENTS, APISERVER, CONTROLLER_MANAGER, SCHEDULER, and WORKLOADS.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#enable_components ContainerCluster#enable_components}
	EnableComponents *[]*string `field:"required" json:"enableComponents" yaml:"enableComponents"`
}

