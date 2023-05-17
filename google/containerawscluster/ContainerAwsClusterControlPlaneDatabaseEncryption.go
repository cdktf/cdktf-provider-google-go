package containerawscluster


type ContainerAwsClusterControlPlaneDatabaseEncryption struct {
	// The ARN of the AWS KMS key used to encrypt cluster secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_aws_cluster#kms_key_arn ContainerAwsCluster#kms_key_arn}
	KmsKeyArn *string `field:"required" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

