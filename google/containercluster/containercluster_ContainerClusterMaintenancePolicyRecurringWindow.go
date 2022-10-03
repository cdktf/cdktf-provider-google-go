package containercluster


type ContainerClusterMaintenancePolicyRecurringWindow struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#end_time ContainerCluster#end_time}.
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#recurrence ContainerCluster#recurrence}.
	Recurrence *string `field:"required" json:"recurrence" yaml:"recurrence"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#start_time ContainerCluster#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

