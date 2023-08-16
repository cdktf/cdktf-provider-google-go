package certificatemanagercertificate


type CertificateManagerCertificateSelfManaged struct {
	// **Deprecated** The certificate chain in PEM-encoded form.
	//
	// Leaf certificate comes first, followed by intermediate ones if any.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/certificate_manager_certificate#certificate_pem CertificateManagerCertificate#certificate_pem}
	CertificatePem *string `field:"optional" json:"certificatePem" yaml:"certificatePem"`
	// The certificate chain in PEM-encoded form.
	//
	// Leaf certificate comes first, followed by intermediate ones if any.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/certificate_manager_certificate#pem_certificate CertificateManagerCertificate#pem_certificate}
	PemCertificate *string `field:"optional" json:"pemCertificate" yaml:"pemCertificate"`
	// The private key of the leaf certificate in PEM-encoded form.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/certificate_manager_certificate#pem_private_key CertificateManagerCertificate#pem_private_key}
	PemPrivateKey *string `field:"optional" json:"pemPrivateKey" yaml:"pemPrivateKey"`
	// **Deprecated** The private key of the leaf certificate in PEM-encoded form.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/certificate_manager_certificate#private_key_pem CertificateManagerCertificate#private_key_pem}
	PrivateKeyPem *string `field:"optional" json:"privateKeyPem" yaml:"privateKeyPem"`
}

