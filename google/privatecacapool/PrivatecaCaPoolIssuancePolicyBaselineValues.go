package privatecacapool


type PrivatecaCaPoolIssuancePolicyBaselineValues struct {
	// ca_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_ca_pool#ca_options PrivatecaCaPool#ca_options}
	CaOptions *PrivatecaCaPoolIssuancePolicyBaselineValuesCaOptions `field:"required" json:"caOptions" yaml:"caOptions"`
	// key_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_ca_pool#key_usage PrivatecaCaPool#key_usage}
	KeyUsage *PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsage `field:"required" json:"keyUsage" yaml:"keyUsage"`
	// additional_extensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_ca_pool#additional_extensions PrivatecaCaPool#additional_extensions}
	AdditionalExtensions interface{} `field:"optional" json:"additionalExtensions" yaml:"additionalExtensions"`
	// Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_ca_pool#aia_ocsp_servers PrivatecaCaPool#aia_ocsp_servers}
	AiaOcspServers *[]*string `field:"optional" json:"aiaOcspServers" yaml:"aiaOcspServers"`
	// name_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_ca_pool#name_constraints PrivatecaCaPool#name_constraints}
	NameConstraints *PrivatecaCaPoolIssuancePolicyBaselineValuesNameConstraints `field:"optional" json:"nameConstraints" yaml:"nameConstraints"`
	// policy_ids block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_ca_pool#policy_ids PrivatecaCaPool#policy_ids}
	PolicyIds interface{} `field:"optional" json:"policyIds" yaml:"policyIds"`
}

