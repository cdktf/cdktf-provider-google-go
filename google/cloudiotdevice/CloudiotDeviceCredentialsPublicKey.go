package cloudiotdevice


type CloudiotDeviceCredentialsPublicKey struct {
	// The format of the key. Possible values: ["RSA_PEM", "RSA_X509_PEM", "ES256_PEM", "ES256_X509_PEM"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudiot_device#format CloudiotDevice#format}
	Format *string `field:"required" json:"format" yaml:"format"`
	// The key data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudiot_device#key CloudiotDevice#key}
	Key *string `field:"required" json:"key" yaml:"key"`
}

