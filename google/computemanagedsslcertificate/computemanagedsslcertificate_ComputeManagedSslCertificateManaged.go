package computemanagedsslcertificate


type ComputeManagedSslCertificateManaged struct {
	// Domains for which a managed SSL certificate will be valid.
	//
	// Currently,
	// there can be up to 100 domains in this list.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_managed_ssl_certificate#domains ComputeManagedSslCertificate#domains}
	Domains *[]*string `field:"required" json:"domains" yaml:"domains"`
}

