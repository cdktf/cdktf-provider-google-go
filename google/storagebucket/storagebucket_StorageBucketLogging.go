package storagebucket


type StorageBucketLogging struct {
	// The bucket that will receive log objects.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket#log_bucket StorageBucket#log_bucket}
	LogBucket *string `field:"required" json:"logBucket" yaml:"logBucket"`
	// The object prefix for log objects.
	//
	// If it's not provided, by default Google Cloud Storage sets this to this bucket's name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket#log_object_prefix StorageBucket#log_object_prefix}
	LogObjectPrefix *string `field:"optional" json:"logObjectPrefix" yaml:"logObjectPrefix"`
}

