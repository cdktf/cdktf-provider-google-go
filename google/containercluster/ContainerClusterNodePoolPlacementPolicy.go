package containercluster


type ContainerClusterNodePoolPlacementPolicy struct {
	// Type defines the type of placement policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_cluster#type ContainerCluster#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

