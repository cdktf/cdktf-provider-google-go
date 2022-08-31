// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type StorageBucketRetentionPolicy struct {
	// The period of time, in seconds, that objects in the bucket must be retained and cannot be deleted, overwritten, or archived.
	//
	// The value must be less than 3,155,760,000 seconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket#retention_period StorageBucket#retention_period}
	RetentionPeriod *float64 `field:"required" json:"retentionPeriod" yaml:"retentionPeriod"`
	// If set to true, the bucket will be locked and permanently restrict edits to the bucket's retention policy.
	//
	// Caution: Locking a bucket is an irreversible action.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket#is_locked StorageBucket#is_locked}
	IsLocked interface{} `field:"optional" json:"isLocked" yaml:"isLocked"`
}

