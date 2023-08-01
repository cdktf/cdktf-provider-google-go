package certificatemanagercertificate


type CertificateManagerCertificateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/certificate_manager_certificate#create CertificateManagerCertificate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/certificate_manager_certificate#delete CertificateManagerCertificate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/certificate_manager_certificate#update CertificateManagerCertificate#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

