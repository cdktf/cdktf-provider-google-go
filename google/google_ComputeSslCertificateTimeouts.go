// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeSslCertificateTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_ssl_certificate#create ComputeSslCertificate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_ssl_certificate#delete ComputeSslCertificate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

