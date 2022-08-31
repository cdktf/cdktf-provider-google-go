// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type PrivatecaCertificateAuthorityConfigSubjectConfig struct {
	// subject block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_authority#subject PrivatecaCertificateAuthority#subject}
	Subject *PrivatecaCertificateAuthorityConfigSubjectConfigSubject `field:"required" json:"subject" yaml:"subject"`
	// subject_alt_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_authority#subject_alt_name PrivatecaCertificateAuthority#subject_alt_name}
	SubjectAltName *PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltName `field:"optional" json:"subjectAltName" yaml:"subjectAltName"`
}

