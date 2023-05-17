package containerawsnodepool


type ContainerAwsNodePoolConfigSshConfig struct {
	// The name of the EC2 key pair used to login into cluster machines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_aws_node_pool#ec2_key_pair ContainerAwsNodePool#ec2_key_pair}
	Ec2KeyPair *string `field:"required" json:"ec2KeyPair" yaml:"ec2KeyPair"`
}

