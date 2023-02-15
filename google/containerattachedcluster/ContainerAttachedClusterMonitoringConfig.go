package containerattachedcluster


type ContainerAttachedClusterMonitoringConfig struct {
	// managed_prometheus_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_attached_cluster#managed_prometheus_config ContainerAttachedCluster#managed_prometheus_config}
	ManagedPrometheusConfig *ContainerAttachedClusterMonitoringConfigManagedPrometheusConfig `field:"optional" json:"managedPrometheusConfig" yaml:"managedPrometheusConfig"`
}

