package privatecacapool


type PrivatecaCaPoolIssuancePolicy struct {
	// allowed_issuance_modes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_ca_pool#allowed_issuance_modes PrivatecaCaPool#allowed_issuance_modes}
	AllowedIssuanceModes *PrivatecaCaPoolIssuancePolicyAllowedIssuanceModes `field:"optional" json:"allowedIssuanceModes" yaml:"allowedIssuanceModes"`
	// allowed_key_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_ca_pool#allowed_key_types PrivatecaCaPool#allowed_key_types}
	AllowedKeyTypes interface{} `field:"optional" json:"allowedKeyTypes" yaml:"allowedKeyTypes"`
	// baseline_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_ca_pool#baseline_values PrivatecaCaPool#baseline_values}
	BaselineValues *PrivatecaCaPoolIssuancePolicyBaselineValues `field:"optional" json:"baselineValues" yaml:"baselineValues"`
	// identity_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_ca_pool#identity_constraints PrivatecaCaPool#identity_constraints}
	IdentityConstraints *PrivatecaCaPoolIssuancePolicyIdentityConstraints `field:"optional" json:"identityConstraints" yaml:"identityConstraints"`
	// The maximum lifetime allowed for issued Certificates.
	//
	// Note that if the issuing CertificateAuthority
	// expires before a Certificate's requested maximumLifetime, the effective lifetime will be explicitly truncated to match it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_ca_pool#maximum_lifetime PrivatecaCaPool#maximum_lifetime}
	MaximumLifetime *string `field:"optional" json:"maximumLifetime" yaml:"maximumLifetime"`
}

