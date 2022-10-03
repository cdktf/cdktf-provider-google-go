package cloudiotregistry


type CloudiotRegistryTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudiot_registry#create CloudiotRegistry#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudiot_registry#delete CloudiotRegistry#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudiot_registry#update CloudiotRegistry#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

