package containercluster


type ContainerClusterMaintenancePolicyDailyMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_cluster#start_time ContainerCluster#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

