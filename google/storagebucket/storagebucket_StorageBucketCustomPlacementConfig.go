package storagebucket


type StorageBucketCustomPlacementConfig struct {
	// The list of individual regions that comprise a dual-region bucket.
	//
	// See the docs for a list of acceptable regions. Note: If any of the data_locations changes, it will recreate the bucket.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket#data_locations StorageBucket#data_locations}
	DataLocations *[]*string `field:"required" json:"dataLocations" yaml:"dataLocations"`
}

