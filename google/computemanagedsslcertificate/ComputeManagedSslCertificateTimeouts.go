package computemanagedsslcertificate


type ComputeManagedSslCertificateTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_managed_ssl_certificate#create ComputeManagedSslCertificate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_managed_ssl_certificate#delete ComputeManagedSslCertificate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

