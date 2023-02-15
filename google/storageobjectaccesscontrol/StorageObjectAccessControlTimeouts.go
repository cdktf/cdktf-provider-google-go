package storageobjectaccesscontrol


type StorageObjectAccessControlTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_object_access_control#create StorageObjectAccessControl#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_object_access_control#delete StorageObjectAccessControl#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_object_access_control#update StorageObjectAccessControl#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

