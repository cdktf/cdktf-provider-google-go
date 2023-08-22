package privatecacapool


type PrivatecaCaPoolIssuancePolicyAllowedIssuanceModes struct {
	// When true, allows callers to create Certificates by specifying a CertificateConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/privateca_ca_pool#allow_config_based_issuance PrivatecaCaPool#allow_config_based_issuance}
	AllowConfigBasedIssuance interface{} `field:"required" json:"allowConfigBasedIssuance" yaml:"allowConfigBasedIssuance"`
	// When true, allows callers to create Certificates by specifying a CSR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/privateca_ca_pool#allow_csr_based_issuance PrivatecaCaPool#allow_csr_based_issuance}
	AllowCsrBasedIssuance interface{} `field:"required" json:"allowCsrBasedIssuance" yaml:"allowCsrBasedIssuance"`
}

