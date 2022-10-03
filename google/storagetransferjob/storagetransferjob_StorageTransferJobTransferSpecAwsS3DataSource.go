package storagetransferjob


type StorageTransferJobTransferSpecAwsS3DataSource struct {
	// S3 Bucket name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#bucket_name StorageTransferJob#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// aws_access_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#aws_access_key StorageTransferJob#aws_access_key}
	AwsAccessKey *StorageTransferJobTransferSpecAwsS3DataSourceAwsAccessKey `field:"optional" json:"awsAccessKey" yaml:"awsAccessKey"`
	// The Amazon Resource Name (ARN) of the role to support temporary credentials via 'AssumeRoleWithWebIdentity'.
	//
	// For more information about ARNs, see [IAM ARNs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns). When a role ARN is provided, Transfer Service fetches temporary credentials for the session using a 'AssumeRoleWithWebIdentity' call for the provided role using the [GoogleServiceAccount][] for this project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#role_arn StorageTransferJob#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

