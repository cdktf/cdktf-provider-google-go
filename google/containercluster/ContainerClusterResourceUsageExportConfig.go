package containercluster


type ContainerClusterResourceUsageExportConfig struct {
	// bigquery_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_cluster#bigquery_destination ContainerCluster#bigquery_destination}
	BigqueryDestination *ContainerClusterResourceUsageExportConfigBigqueryDestination `field:"required" json:"bigqueryDestination" yaml:"bigqueryDestination"`
	// Whether to enable network egress metering for this cluster.
	//
	// If enabled, a daemonset will be created in the cluster to meter network egress traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_cluster#enable_network_egress_metering ContainerCluster#enable_network_egress_metering}
	EnableNetworkEgressMetering interface{} `field:"optional" json:"enableNetworkEgressMetering" yaml:"enableNetworkEgressMetering"`
	// Whether to enable resource consumption metering on this cluster.
	//
	// When enabled, a table will be created in the resource export BigQuery dataset to store resource consumption data. The resulting table can be joined with the resource usage table or with BigQuery billing export. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_cluster#enable_resource_consumption_metering ContainerCluster#enable_resource_consumption_metering}
	EnableResourceConsumptionMetering interface{} `field:"optional" json:"enableResourceConsumptionMetering" yaml:"enableResourceConsumptionMetering"`
}

