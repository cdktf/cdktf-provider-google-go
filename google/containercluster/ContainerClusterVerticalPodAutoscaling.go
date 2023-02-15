package containercluster


type ContainerClusterVerticalPodAutoscaling struct {
	// Enables vertical pod autoscaling.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

