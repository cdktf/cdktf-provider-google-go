package privatecacertificatetemplate


type PrivatecaCertificateTemplatePredefinedValues struct {
	// additional_extensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_certificate_template#additional_extensions PrivatecaCertificateTemplate#additional_extensions}
	AdditionalExtensions interface{} `field:"optional" json:"additionalExtensions" yaml:"additionalExtensions"`
	// Optional.
	//
	// Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_certificate_template#aia_ocsp_servers PrivatecaCertificateTemplate#aia_ocsp_servers}
	AiaOcspServers *[]*string `field:"optional" json:"aiaOcspServers" yaml:"aiaOcspServers"`
	// ca_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_certificate_template#ca_options PrivatecaCertificateTemplate#ca_options}
	CaOptions *PrivatecaCertificateTemplatePredefinedValuesCaOptions `field:"optional" json:"caOptions" yaml:"caOptions"`
	// key_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_certificate_template#key_usage PrivatecaCertificateTemplate#key_usage}
	KeyUsage *PrivatecaCertificateTemplatePredefinedValuesKeyUsage `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// policy_ids block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_certificate_template#policy_ids PrivatecaCertificateTemplate#policy_ids}
	PolicyIds interface{} `field:"optional" json:"policyIds" yaml:"policyIds"`
}

