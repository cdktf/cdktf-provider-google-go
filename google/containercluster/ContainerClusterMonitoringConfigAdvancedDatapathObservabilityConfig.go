package containercluster


type ContainerClusterMonitoringConfigAdvancedDatapathObservabilityConfig struct {
	// Whether or not the advanced datapath metrics are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#enable_metrics ContainerCluster#enable_metrics}
	EnableMetrics interface{} `field:"required" json:"enableMetrics" yaml:"enableMetrics"`
	// Mode used to make Relay available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#relay_mode ContainerCluster#relay_mode}
	RelayMode *string `field:"optional" json:"relayMode" yaml:"relayMode"`
}

