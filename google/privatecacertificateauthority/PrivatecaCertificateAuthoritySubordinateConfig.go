package privatecacertificateauthority


type PrivatecaCertificateAuthoritySubordinateConfig struct {
	// This can refer to a CertificateAuthority that was used to create a subordinate CertificateAuthority.
	//
	// This field is used for information
	// and usability purposes only. The resource name is in the format
	// 'projects/*\/locations/*\/caPools/*\/certificateAuthorities/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#certificate_authority PrivatecaCertificateAuthority#certificate_authority}
	CertificateAuthority *string `field:"optional" json:"certificateAuthority" yaml:"certificateAuthority"`
	// pem_issuer_chain block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#pem_issuer_chain PrivatecaCertificateAuthority#pem_issuer_chain}
	PemIssuerChain *PrivatecaCertificateAuthoritySubordinateConfigPemIssuerChain `field:"optional" json:"pemIssuerChain" yaml:"pemIssuerChain"`
}

