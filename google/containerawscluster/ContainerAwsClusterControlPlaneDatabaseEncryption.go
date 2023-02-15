package containerawscluster


type ContainerAwsClusterControlPlaneDatabaseEncryption struct {
	// The ARN of the AWS KMS key used to encrypt cluster secrets.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_aws_cluster#kms_key_arn ContainerAwsCluster#kms_key_arn}
	KmsKeyArn *string `field:"required" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

