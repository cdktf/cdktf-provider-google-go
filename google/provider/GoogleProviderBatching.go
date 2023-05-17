package provider


type GoogleProviderBatching struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs#enable_batching GoogleProvider#enable_batching}.
	EnableBatching interface{} `field:"optional" json:"enableBatching" yaml:"enableBatching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs#send_after GoogleProvider#send_after}.
	SendAfter *string `field:"optional" json:"sendAfter" yaml:"sendAfter"`
}

