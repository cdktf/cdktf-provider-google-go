package provider


type GoogleProviderBatching struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#enable_batching GoogleProvider#enable_batching}.
	EnableBatching interface{} `field:"optional" json:"enableBatching" yaml:"enableBatching"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#send_after GoogleProvider#send_after}.
	SendAfter *string `field:"optional" json:"sendAfter" yaml:"sendAfter"`
}

