// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type PrivatecaCertificateConfigX509ConfigKeyUsage struct {
	// base_key_usage block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate#base_key_usage PrivatecaCertificate#base_key_usage}
	BaseKeyUsage *PrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsage `field:"required" json:"baseKeyUsage" yaml:"baseKeyUsage"`
	// extended_key_usage block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate#extended_key_usage PrivatecaCertificate#extended_key_usage}
	ExtendedKeyUsage *PrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage `field:"required" json:"extendedKeyUsage" yaml:"extendedKeyUsage"`
	// unknown_extended_key_usages block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate#unknown_extended_key_usages PrivatecaCertificate#unknown_extended_key_usages}
	UnknownExtendedKeyUsages interface{} `field:"optional" json:"unknownExtendedKeyUsages" yaml:"unknownExtendedKeyUsages"`
}

