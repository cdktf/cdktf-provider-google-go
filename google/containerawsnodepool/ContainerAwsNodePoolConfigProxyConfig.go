package containerawsnodepool


type ContainerAwsNodePoolConfigProxyConfig struct {
	// The ARN of the AWS Secret Manager secret that contains the HTTP(S) proxy configuration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_aws_node_pool#secret_arn ContainerAwsNodePool#secret_arn}
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The version string of the AWS Secret Manager secret that contains the HTTP(S) proxy configuration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_aws_node_pool#secret_version ContainerAwsNodePool#secret_version}
	SecretVersion *string `field:"required" json:"secretVersion" yaml:"secretVersion"`
}

