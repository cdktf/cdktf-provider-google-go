// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagetransferjob


type StorageTransferJobTransferSpecAwsS3DataSource struct {
	// S3 Bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_transfer_job#bucket_name StorageTransferJob#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// aws_access_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_transfer_job#aws_access_key StorageTransferJob#aws_access_key}
	AwsAccessKey *StorageTransferJobTransferSpecAwsS3DataSourceAwsAccessKey `field:"optional" json:"awsAccessKey" yaml:"awsAccessKey"`
	// The CloudFront distribution domain name pointing to this bucket, to use when fetching.
	//
	// See [Transfer from S3 via CloudFront](https://cloud.google.com/storage-transfer/docs/s3-cloudfront) for more information. Format: https://{id}.cloudfront.net or any valid custom domain. Must begin with https://.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_transfer_job#cloudfront_domain StorageTransferJob#cloudfront_domain}
	CloudfrontDomain *string `field:"optional" json:"cloudfrontDomain" yaml:"cloudfrontDomain"`
	// Egress bytes over a Google-managed private network. This network is shared between other users of Storage Transfer Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_transfer_job#managed_private_network StorageTransferJob#managed_private_network}
	ManagedPrivateNetwork interface{} `field:"optional" json:"managedPrivateNetwork" yaml:"managedPrivateNetwork"`
	// S3 Bucket path in bucket to transfer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_transfer_job#path StorageTransferJob#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The Amazon Resource Name (ARN) of the role to support temporary credentials via 'AssumeRoleWithWebIdentity'.
	//
	// For more information about ARNs, see [IAM ARNs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns). When a role ARN is provided, Transfer Service fetches temporary credentials for the session using a 'AssumeRoleWithWebIdentity' call for the provided role using the [GoogleServiceAccount][] for this project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_transfer_job#role_arn StorageTransferJob#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

