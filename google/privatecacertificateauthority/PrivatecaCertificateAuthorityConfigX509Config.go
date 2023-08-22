package privatecacertificateauthority


type PrivatecaCertificateAuthorityConfigX509Config struct {
	// ca_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/privateca_certificate_authority#ca_options PrivatecaCertificateAuthority#ca_options}
	CaOptions *PrivatecaCertificateAuthorityConfigX509ConfigCaOptions `field:"required" json:"caOptions" yaml:"caOptions"`
	// key_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/privateca_certificate_authority#key_usage PrivatecaCertificateAuthority#key_usage}
	KeyUsage *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsage `field:"required" json:"keyUsage" yaml:"keyUsage"`
	// additional_extensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/privateca_certificate_authority#additional_extensions PrivatecaCertificateAuthority#additional_extensions}
	AdditionalExtensions interface{} `field:"optional" json:"additionalExtensions" yaml:"additionalExtensions"`
	// Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/privateca_certificate_authority#aia_ocsp_servers PrivatecaCertificateAuthority#aia_ocsp_servers}
	AiaOcspServers *[]*string `field:"optional" json:"aiaOcspServers" yaml:"aiaOcspServers"`
	// name_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/privateca_certificate_authority#name_constraints PrivatecaCertificateAuthority#name_constraints}
	NameConstraints *PrivatecaCertificateAuthorityConfigX509ConfigNameConstraints `field:"optional" json:"nameConstraints" yaml:"nameConstraints"`
	// policy_ids block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/privateca_certificate_authority#policy_ids PrivatecaCertificateAuthority#policy_ids}
	PolicyIds interface{} `field:"optional" json:"policyIds" yaml:"policyIds"`
}

