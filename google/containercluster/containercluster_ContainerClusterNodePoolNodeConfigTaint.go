package containercluster


type ContainerClusterNodePoolNodeConfigTaint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#effect ContainerCluster#effect}.
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#key ContainerCluster#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#value ContainerCluster#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

