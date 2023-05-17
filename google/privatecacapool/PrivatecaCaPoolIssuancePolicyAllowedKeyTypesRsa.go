package privatecacapool


type PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa struct {
	// The maximum allowed RSA modulus size, in bits.
	//
	// If this is not set, or if set to zero, the
	// service will not enforce an explicit upper bound on RSA modulus sizes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_ca_pool#max_modulus_size PrivatecaCaPool#max_modulus_size}
	MaxModulusSize *string `field:"optional" json:"maxModulusSize" yaml:"maxModulusSize"`
	// The minimum allowed RSA modulus size, in bits.
	//
	// If this is not set, or if set to zero, the
	// service-level min RSA modulus size will continue to apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_ca_pool#min_modulus_size PrivatecaCaPool#min_modulus_size}
	MinModulusSize *string `field:"optional" json:"minModulusSize" yaml:"minModulusSize"`
}

