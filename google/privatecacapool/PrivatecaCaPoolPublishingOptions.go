package privatecacapool


type PrivatecaCaPoolPublishingOptions struct {
	// When true, publishes each CertificateAuthority's CA certificate and includes its URL in the "Authority Information Access" X.509 extension in all issued Certificates. If this is false, the CA certificate will not be published and the corresponding X.509 extension will not be written in issued certificates.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_ca_pool#publish_ca_cert PrivatecaCaPool#publish_ca_cert}
	PublishCaCert interface{} `field:"required" json:"publishCaCert" yaml:"publishCaCert"`
	// When true, publishes each CertificateAuthority's CRL and includes its URL in the "CRL Distribution Points" X.509 extension in all issued Certificates. If this is false, CRLs will not be published and the corresponding X.509 extension will not be written in issued certificates. CRLs will expire 7 days from their creation. However, we will rebuild daily. CRLs are also rebuilt shortly after a certificate is revoked.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_ca_pool#publish_crl PrivatecaCaPool#publish_crl}
	PublishCrl interface{} `field:"required" json:"publishCrl" yaml:"publishCrl"`
}

