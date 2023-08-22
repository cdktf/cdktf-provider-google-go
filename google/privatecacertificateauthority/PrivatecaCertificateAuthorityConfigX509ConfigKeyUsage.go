package privatecacertificateauthority


type PrivatecaCertificateAuthorityConfigX509ConfigKeyUsage struct {
	// base_key_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/privateca_certificate_authority#base_key_usage PrivatecaCertificateAuthority#base_key_usage}
	BaseKeyUsage *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage `field:"required" json:"baseKeyUsage" yaml:"baseKeyUsage"`
	// extended_key_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/privateca_certificate_authority#extended_key_usage PrivatecaCertificateAuthority#extended_key_usage}
	ExtendedKeyUsage *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage `field:"required" json:"extendedKeyUsage" yaml:"extendedKeyUsage"`
	// unknown_extended_key_usages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/privateca_certificate_authority#unknown_extended_key_usages PrivatecaCertificateAuthority#unknown_extended_key_usages}
	UnknownExtendedKeyUsages interface{} `field:"optional" json:"unknownExtendedKeyUsages" yaml:"unknownExtendedKeyUsages"`
}

