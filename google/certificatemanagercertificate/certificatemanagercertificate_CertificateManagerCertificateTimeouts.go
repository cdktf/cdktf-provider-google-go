package certificatemanagercertificate


type CertificateManagerCertificateTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/certificate_manager_certificate#create CertificateManagerCertificate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/certificate_manager_certificate#delete CertificateManagerCertificate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/certificate_manager_certificate#update CertificateManagerCertificate#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

