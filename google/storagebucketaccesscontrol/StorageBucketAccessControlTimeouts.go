package storagebucketaccesscontrol


type StorageBucketAccessControlTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/storage_bucket_access_control#create StorageBucketAccessControl#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/storage_bucket_access_control#delete StorageBucketAccessControl#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/storage_bucket_access_control#update StorageBucketAccessControl#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

