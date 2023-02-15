package certificatemanagerdnsauthorization


type CertificateManagerDnsAuthorizationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/certificate_manager_dns_authorization#create CertificateManagerDnsAuthorization#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/certificate_manager_dns_authorization#delete CertificateManagerDnsAuthorization#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/certificate_manager_dns_authorization#update CertificateManagerDnsAuthorization#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

