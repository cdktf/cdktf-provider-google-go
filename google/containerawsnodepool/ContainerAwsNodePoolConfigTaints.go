package containerawsnodepool


type ContainerAwsNodePoolConfigTaints struct {
	// The taint effect. Possible values: EFFECT_UNSPECIFIED, NO_SCHEDULE, PREFER_NO_SCHEDULE, NO_EXECUTE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_aws_node_pool#effect ContainerAwsNodePool#effect}
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Key for the taint.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_aws_node_pool#key ContainerAwsNodePool#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value for the taint.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_aws_node_pool#value ContainerAwsNodePool#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

