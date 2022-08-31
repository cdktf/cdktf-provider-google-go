// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type StorageBucketVersioning struct {
	// While set to true, versioning is fully enabled for this bucket.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket#enabled StorageBucket#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

