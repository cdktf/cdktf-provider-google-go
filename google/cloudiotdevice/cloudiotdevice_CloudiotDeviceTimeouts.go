package cloudiotdevice


type CloudiotDeviceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudiot_device#create CloudiotDevice#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudiot_device#delete CloudiotDevice#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudiot_device#update CloudiotDevice#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

