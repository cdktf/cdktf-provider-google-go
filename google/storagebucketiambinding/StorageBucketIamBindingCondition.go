package storagebucketiambinding


type StorageBucketIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket_iam_binding#expression StorageBucketIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket_iam_binding#title StorageBucketIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket_iam_binding#description StorageBucketIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

