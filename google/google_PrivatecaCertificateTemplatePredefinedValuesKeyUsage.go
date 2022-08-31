// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type PrivatecaCertificateTemplatePredefinedValuesKeyUsage struct {
	// base_key_usage block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_template#base_key_usage PrivatecaCertificateTemplate#base_key_usage}
	BaseKeyUsage *PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage `field:"optional" json:"baseKeyUsage" yaml:"baseKeyUsage"`
	// extended_key_usage block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_template#extended_key_usage PrivatecaCertificateTemplate#extended_key_usage}
	ExtendedKeyUsage *PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage `field:"optional" json:"extendedKeyUsage" yaml:"extendedKeyUsage"`
	// unknown_extended_key_usages block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_template#unknown_extended_key_usages PrivatecaCertificateTemplate#unknown_extended_key_usages}
	UnknownExtendedKeyUsages interface{} `field:"optional" json:"unknownExtendedKeyUsages" yaml:"unknownExtendedKeyUsages"`
}

