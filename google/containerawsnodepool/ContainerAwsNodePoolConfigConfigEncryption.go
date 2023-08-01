package containerawsnodepool


type ContainerAwsNodePoolConfigConfigEncryption struct {
	// The ARN of the AWS KMS key used to encrypt node pool configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_aws_node_pool#kms_key_arn ContainerAwsNodePool#kms_key_arn}
	KmsKeyArn *string `field:"required" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

