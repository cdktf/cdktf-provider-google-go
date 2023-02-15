package storagebucket


type StorageBucketTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket#create StorageBucket#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket#read StorageBucket#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket#update StorageBucket#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

