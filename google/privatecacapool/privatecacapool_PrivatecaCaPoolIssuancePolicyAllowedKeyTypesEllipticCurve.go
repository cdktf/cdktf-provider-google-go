package privatecacapool


type PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve struct {
	// The algorithm used. Possible values: ["ECDSA_P256", "ECDSA_P384", "EDDSA_25519"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_ca_pool#signature_algorithm PrivatecaCaPool#signature_algorithm}
	SignatureAlgorithm *string `field:"required" json:"signatureAlgorithm" yaml:"signatureAlgorithm"`
}

