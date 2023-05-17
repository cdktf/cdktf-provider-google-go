package storagedefaultobjectaccesscontrol


type StorageDefaultObjectAccessControlTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/storage_default_object_access_control#create StorageDefaultObjectAccessControl#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/storage_default_object_access_control#delete StorageDefaultObjectAccessControl#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/storage_default_object_access_control#update StorageDefaultObjectAccessControl#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

