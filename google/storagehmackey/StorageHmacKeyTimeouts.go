package storagehmackey


type StorageHmacKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_hmac_key#create StorageHmacKey#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_hmac_key#delete StorageHmacKey#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_hmac_key#update StorageHmacKey#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

