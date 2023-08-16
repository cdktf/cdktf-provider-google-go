package privatecacapool


type PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve struct {
	// The algorithm used. Possible values: ["ECDSA_P256", "ECDSA_P384", "EDDSA_25519"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_ca_pool#signature_algorithm PrivatecaCaPool#signature_algorithm}
	SignatureAlgorithm *string `field:"required" json:"signatureAlgorithm" yaml:"signatureAlgorithm"`
}

