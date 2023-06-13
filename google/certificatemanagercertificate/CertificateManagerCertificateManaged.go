package certificatemanagercertificate


type CertificateManagerCertificateManaged struct {
	// Authorizations that will be used for performing domain authorization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/certificate_manager_certificate#dns_authorizations CertificateManagerCertificate#dns_authorizations}
	DnsAuthorizations *[]*string `field:"optional" json:"dnsAuthorizations" yaml:"dnsAuthorizations"`
	// The domains for which a managed SSL certificate will be generated. Wildcard domains are only supported with DNS challenge resolution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/certificate_manager_certificate#domains CertificateManagerCertificate#domains}
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
}

