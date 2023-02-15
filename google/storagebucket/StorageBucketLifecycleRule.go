package storagebucket


type StorageBucketLifecycleRule struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket#action StorageBucket#action}
	Action *StorageBucketLifecycleRuleAction `field:"required" json:"action" yaml:"action"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket#condition StorageBucket#condition}
	Condition *StorageBucketLifecycleRuleCondition `field:"required" json:"condition" yaml:"condition"`
}

