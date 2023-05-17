package containerawscluster


type ContainerAwsClusterControlPlaneAwsServicesAuthentication struct {
	// The Amazon Resource Name (ARN) of the role that the Anthos Multi-Cloud API will assume when managing AWS resources on your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_aws_cluster#role_arn ContainerAwsCluster#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Optional. An identifier for the assumed role session. When unspecified, it defaults to `multicloud-service-agent`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_aws_cluster#role_session_name ContainerAwsCluster#role_session_name}
	RoleSessionName *string `field:"optional" json:"roleSessionName" yaml:"roleSessionName"`
}

