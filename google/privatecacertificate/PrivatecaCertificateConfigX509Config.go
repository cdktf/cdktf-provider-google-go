package privatecacertificate


type PrivatecaCertificateConfigX509Config struct {
	// key_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate#key_usage PrivatecaCertificate#key_usage}
	KeyUsage *PrivatecaCertificateConfigX509ConfigKeyUsage `field:"required" json:"keyUsage" yaml:"keyUsage"`
	// additional_extensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate#additional_extensions PrivatecaCertificate#additional_extensions}
	AdditionalExtensions interface{} `field:"optional" json:"additionalExtensions" yaml:"additionalExtensions"`
	// Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate#aia_ocsp_servers PrivatecaCertificate#aia_ocsp_servers}
	AiaOcspServers *[]*string `field:"optional" json:"aiaOcspServers" yaml:"aiaOcspServers"`
	// ca_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate#ca_options PrivatecaCertificate#ca_options}
	CaOptions *PrivatecaCertificateConfigX509ConfigCaOptions `field:"optional" json:"caOptions" yaml:"caOptions"`
	// name_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate#name_constraints PrivatecaCertificate#name_constraints}
	NameConstraints *PrivatecaCertificateConfigX509ConfigNameConstraints `field:"optional" json:"nameConstraints" yaml:"nameConstraints"`
	// policy_ids block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate#policy_ids PrivatecaCertificate#policy_ids}
	PolicyIds interface{} `field:"optional" json:"policyIds" yaml:"policyIds"`
}

