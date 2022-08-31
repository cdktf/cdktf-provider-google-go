// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ContainerAwsClusterControlPlaneConfigEncryption struct {
	// The ARN of the AWS KMS key used to encrypt cluster configuration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_aws_cluster#kms_key_arn ContainerAwsCluster#kms_key_arn}
	KmsKeyArn *string `field:"required" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

