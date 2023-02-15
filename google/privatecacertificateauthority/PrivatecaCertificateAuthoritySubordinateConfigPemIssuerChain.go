package privatecacertificateauthority


type PrivatecaCertificateAuthoritySubordinateConfigPemIssuerChain struct {
	// Expected to be in leaf-to-root order according to RFC 5246.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_authority#pem_certificates PrivatecaCertificateAuthority#pem_certificates}
	PemCertificates *[]*string `field:"optional" json:"pemCertificates" yaml:"pemCertificates"`
}

