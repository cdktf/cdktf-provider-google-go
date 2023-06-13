package privatecacapool


type PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage struct {
	// The key may be used to sign certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_ca_pool#cert_sign PrivatecaCaPool#cert_sign}
	CertSign interface{} `field:"optional" json:"certSign" yaml:"certSign"`
	// The key may be used for cryptographic commitments. Note that this may also be referred to as "non-repudiation".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_ca_pool#content_commitment PrivatecaCaPool#content_commitment}
	ContentCommitment interface{} `field:"optional" json:"contentCommitment" yaml:"contentCommitment"`
	// The key may be used sign certificate revocation lists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_ca_pool#crl_sign PrivatecaCaPool#crl_sign}
	CrlSign interface{} `field:"optional" json:"crlSign" yaml:"crlSign"`
	// The key may be used to encipher data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_ca_pool#data_encipherment PrivatecaCaPool#data_encipherment}
	DataEncipherment interface{} `field:"optional" json:"dataEncipherment" yaml:"dataEncipherment"`
	// The key may be used to decipher only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_ca_pool#decipher_only PrivatecaCaPool#decipher_only}
	DecipherOnly interface{} `field:"optional" json:"decipherOnly" yaml:"decipherOnly"`
	// The key may be used for digital signatures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_ca_pool#digital_signature PrivatecaCaPool#digital_signature}
	DigitalSignature interface{} `field:"optional" json:"digitalSignature" yaml:"digitalSignature"`
	// The key may be used to encipher only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_ca_pool#encipher_only PrivatecaCaPool#encipher_only}
	EncipherOnly interface{} `field:"optional" json:"encipherOnly" yaml:"encipherOnly"`
	// The key may be used in a key agreement protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_ca_pool#key_agreement PrivatecaCaPool#key_agreement}
	KeyAgreement interface{} `field:"optional" json:"keyAgreement" yaml:"keyAgreement"`
	// The key may be used to encipher other keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_ca_pool#key_encipherment PrivatecaCaPool#key_encipherment}
	KeyEncipherment interface{} `field:"optional" json:"keyEncipherment" yaml:"keyEncipherment"`
}

