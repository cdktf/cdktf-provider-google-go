package certificatemanagercertificate


type CertificateManagerCertificateManaged struct {
	// Authorizations that will be used for performing domain authorization. Either issuanceConfig or dnsAuthorizations should be specificed, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/certificate_manager_certificate#dns_authorizations CertificateManagerCertificate#dns_authorizations}
	DnsAuthorizations *[]*string `field:"optional" json:"dnsAuthorizations" yaml:"dnsAuthorizations"`
	// The domains for which a managed SSL certificate will be generated. Wildcard domains are only supported with DNS challenge resolution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/certificate_manager_certificate#domains CertificateManagerCertificate#domains}
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// The resource name for a CertificateIssuanceConfig used to configure private PKI certificates in the format projects/*\/locations/*\/certificateIssuanceConfigs/*.
	//
	// If this field is not set, the certificates will instead be publicly signed as documented at https://cloud.google.com/load-balancing/docs/ssl-certificates/google-managed-certs#caa.
	// Either issuanceConfig or dnsAuthorizations should be specificed, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/certificate_manager_certificate#issuance_config CertificateManagerCertificate#issuance_config}
	IssuanceConfig *string `field:"optional" json:"issuanceConfig" yaml:"issuanceConfig"`
}

