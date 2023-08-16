package privatecacapool


type PrivatecaCaPoolIssuancePolicyAllowedKeyTypes struct {
	// elliptic_curve block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_ca_pool#elliptic_curve PrivatecaCaPool#elliptic_curve}
	EllipticCurve *PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve `field:"optional" json:"ellipticCurve" yaml:"ellipticCurve"`
	// rsa block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_ca_pool#rsa PrivatecaCaPool#rsa}
	Rsa *PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa `field:"optional" json:"rsa" yaml:"rsa"`
}

