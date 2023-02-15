package containercluster


type ContainerClusterNodePoolPlacementPolicy struct {
	// Type defines the type of placement policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#type ContainerCluster#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

