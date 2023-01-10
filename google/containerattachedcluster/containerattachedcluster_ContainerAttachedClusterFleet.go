package containerattachedcluster


type ContainerAttachedClusterFleet struct {
	// The number of the Fleet host project where this cluster will be registered.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_attached_cluster#project ContainerAttachedCluster#project}
	Project *string `field:"required" json:"project" yaml:"project"`
}

