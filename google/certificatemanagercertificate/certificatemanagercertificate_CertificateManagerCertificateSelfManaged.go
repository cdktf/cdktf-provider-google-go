package certificatemanagercertificate


type CertificateManagerCertificateSelfManaged struct {
	// The certificate chain in PEM-encoded form.
	//
	// Leaf certificate comes first, followed by intermediate ones if any.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/certificate_manager_certificate#certificate_pem CertificateManagerCertificate#certificate_pem}
	CertificatePem *string `field:"required" json:"certificatePem" yaml:"certificatePem"`
	// The private key of the leaf certificate in PEM-encoded form.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/certificate_manager_certificate#private_key_pem CertificateManagerCertificate#private_key_pem}
	PrivateKeyPem *string `field:"required" json:"privateKeyPem" yaml:"privateKeyPem"`
}

