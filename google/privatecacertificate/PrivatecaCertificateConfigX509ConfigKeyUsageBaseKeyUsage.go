package privatecacertificate


type PrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsage struct {
	// The key may be used to sign certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate#cert_sign PrivatecaCertificate#cert_sign}
	CertSign interface{} `field:"optional" json:"certSign" yaml:"certSign"`
	// The key may be used for cryptographic commitments. Note that this may also be referred to as "non-repudiation".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate#content_commitment PrivatecaCertificate#content_commitment}
	ContentCommitment interface{} `field:"optional" json:"contentCommitment" yaml:"contentCommitment"`
	// The key may be used sign certificate revocation lists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate#crl_sign PrivatecaCertificate#crl_sign}
	CrlSign interface{} `field:"optional" json:"crlSign" yaml:"crlSign"`
	// The key may be used to encipher data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate#data_encipherment PrivatecaCertificate#data_encipherment}
	DataEncipherment interface{} `field:"optional" json:"dataEncipherment" yaml:"dataEncipherment"`
	// The key may be used to decipher only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate#decipher_only PrivatecaCertificate#decipher_only}
	DecipherOnly interface{} `field:"optional" json:"decipherOnly" yaml:"decipherOnly"`
	// The key may be used for digital signatures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate#digital_signature PrivatecaCertificate#digital_signature}
	DigitalSignature interface{} `field:"optional" json:"digitalSignature" yaml:"digitalSignature"`
	// The key may be used to encipher only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate#encipher_only PrivatecaCertificate#encipher_only}
	EncipherOnly interface{} `field:"optional" json:"encipherOnly" yaml:"encipherOnly"`
	// The key may be used in a key agreement protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate#key_agreement PrivatecaCertificate#key_agreement}
	KeyAgreement interface{} `field:"optional" json:"keyAgreement" yaml:"keyAgreement"`
	// The key may be used to encipher other keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate#key_encipherment PrivatecaCertificate#key_encipherment}
	KeyEncipherment interface{} `field:"optional" json:"keyEncipherment" yaml:"keyEncipherment"`
}

