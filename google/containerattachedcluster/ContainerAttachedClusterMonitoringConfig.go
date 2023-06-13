package containerattachedcluster


type ContainerAttachedClusterMonitoringConfig struct {
	// managed_prometheus_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_attached_cluster#managed_prometheus_config ContainerAttachedCluster#managed_prometheus_config}
	ManagedPrometheusConfig *ContainerAttachedClusterMonitoringConfigManagedPrometheusConfig `field:"optional" json:"managedPrometheusConfig" yaml:"managedPrometheusConfig"`
}

