package privatecacertificatetemplate


type PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage struct {
	// The key may be used to sign certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate_template#cert_sign PrivatecaCertificateTemplate#cert_sign}
	CertSign interface{} `field:"optional" json:"certSign" yaml:"certSign"`
	// The key may be used for cryptographic commitments. Note that this may also be referred to as "non-repudiation".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate_template#content_commitment PrivatecaCertificateTemplate#content_commitment}
	ContentCommitment interface{} `field:"optional" json:"contentCommitment" yaml:"contentCommitment"`
	// The key may be used sign certificate revocation lists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate_template#crl_sign PrivatecaCertificateTemplate#crl_sign}
	CrlSign interface{} `field:"optional" json:"crlSign" yaml:"crlSign"`
	// The key may be used to encipher data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate_template#data_encipherment PrivatecaCertificateTemplate#data_encipherment}
	DataEncipherment interface{} `field:"optional" json:"dataEncipherment" yaml:"dataEncipherment"`
	// The key may be used to decipher only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate_template#decipher_only PrivatecaCertificateTemplate#decipher_only}
	DecipherOnly interface{} `field:"optional" json:"decipherOnly" yaml:"decipherOnly"`
	// The key may be used for digital signatures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate_template#digital_signature PrivatecaCertificateTemplate#digital_signature}
	DigitalSignature interface{} `field:"optional" json:"digitalSignature" yaml:"digitalSignature"`
	// The key may be used to encipher only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate_template#encipher_only PrivatecaCertificateTemplate#encipher_only}
	EncipherOnly interface{} `field:"optional" json:"encipherOnly" yaml:"encipherOnly"`
	// The key may be used in a key agreement protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate_template#key_agreement PrivatecaCertificateTemplate#key_agreement}
	KeyAgreement interface{} `field:"optional" json:"keyAgreement" yaml:"keyAgreement"`
	// The key may be used to encipher other keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate_template#key_encipherment PrivatecaCertificateTemplate#key_encipherment}
	KeyEncipherment interface{} `field:"optional" json:"keyEncipherment" yaml:"keyEncipherment"`
}

