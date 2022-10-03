package containerawscluster


type ContainerAwsClusterFleet struct {
	// The number of the Fleet host project where this cluster will be registered.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_aws_cluster#project ContainerAwsCluster#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

