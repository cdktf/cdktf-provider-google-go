package containerawscluster


type ContainerAwsClusterControlPlaneConfigEncryption struct {
	// The ARN of the AWS KMS key used to encrypt cluster configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.75.1/docs/resources/container_aws_cluster#kms_key_arn ContainerAwsCluster#kms_key_arn}
	KmsKeyArn *string `field:"required" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

