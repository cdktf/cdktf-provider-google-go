package containerawscluster


type ContainerAwsClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.70.0/docs/resources/container_aws_cluster#create ContainerAwsCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.70.0/docs/resources/container_aws_cluster#delete ContainerAwsCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.70.0/docs/resources/container_aws_cluster#update ContainerAwsCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

