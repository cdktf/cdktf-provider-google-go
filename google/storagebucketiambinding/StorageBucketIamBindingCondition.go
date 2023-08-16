package storagebucketiambinding


type StorageBucketIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/storage_bucket_iam_binding#expression StorageBucketIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/storage_bucket_iam_binding#title StorageBucketIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/storage_bucket_iam_binding#description StorageBucketIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

