package containercluster


type ContainerClusterMonitoringConfig struct {
	// advanced_datapath_observability_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#advanced_datapath_observability_config ContainerCluster#advanced_datapath_observability_config}
	AdvancedDatapathObservabilityConfig interface{} `field:"optional" json:"advancedDatapathObservabilityConfig" yaml:"advancedDatapathObservabilityConfig"`
	// GKE components exposing metrics. Valid values include SYSTEM_COMPONENTS, APISERVER, CONTROLLER_MANAGER, and SCHEDULER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#enable_components ContainerCluster#enable_components}
	EnableComponents *[]*string `field:"optional" json:"enableComponents" yaml:"enableComponents"`
	// managed_prometheus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#managed_prometheus ContainerCluster#managed_prometheus}
	ManagedPrometheus *ContainerClusterMonitoringConfigManagedPrometheus `field:"optional" json:"managedPrometheus" yaml:"managedPrometheus"`
}

