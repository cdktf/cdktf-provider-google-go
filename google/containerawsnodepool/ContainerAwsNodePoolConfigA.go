package containerawsnodepool


type ContainerAwsNodePoolConfigA struct {
	// config_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_aws_node_pool#config_encryption ContainerAwsNodePool#config_encryption}
	ConfigEncryption *ContainerAwsNodePoolConfigConfigEncryption `field:"required" json:"configEncryption" yaml:"configEncryption"`
	// The name of the AWS IAM role assigned to nodes in the pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_aws_node_pool#iam_instance_profile ContainerAwsNodePool#iam_instance_profile}
	IamInstanceProfile *string `field:"required" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// autoscaling_metrics_collection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_aws_node_pool#autoscaling_metrics_collection ContainerAwsNodePool#autoscaling_metrics_collection}
	AutoscalingMetricsCollection *ContainerAwsNodePoolConfigAutoscalingMetricsCollection `field:"optional" json:"autoscalingMetricsCollection" yaml:"autoscalingMetricsCollection"`
	// Optional. The AWS instance type. When unspecified, it defaults to `m5.large`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_aws_node_pool#instance_type ContainerAwsNodePool#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Optional.
	//
	// The initial labels assigned to nodes of this node pool. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_aws_node_pool#labels ContainerAwsNodePool#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// proxy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_aws_node_pool#proxy_config ContainerAwsNodePool#proxy_config}
	ProxyConfig *ContainerAwsNodePoolConfigProxyConfig `field:"optional" json:"proxyConfig" yaml:"proxyConfig"`
	// root_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_aws_node_pool#root_volume ContainerAwsNodePool#root_volume}
	RootVolume *ContainerAwsNodePoolConfigRootVolume `field:"optional" json:"rootVolume" yaml:"rootVolume"`
	// Optional.
	//
	// The IDs of additional security groups to add to nodes in this pool. The manager will automatically create security groups with minimum rules needed for a functioning cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_aws_node_pool#security_group_ids ContainerAwsNodePool#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// ssh_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_aws_node_pool#ssh_config ContainerAwsNodePool#ssh_config}
	SshConfig *ContainerAwsNodePoolConfigSshConfig `field:"optional" json:"sshConfig" yaml:"sshConfig"`
	// Optional.
	//
	// Key/value metadata to assign to each underlying AWS resource. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_aws_node_pool#tags ContainerAwsNodePool#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// taints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_aws_node_pool#taints ContainerAwsNodePool#taints}
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
}

