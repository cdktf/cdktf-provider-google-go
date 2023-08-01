package privatecacertificateauthority


type PrivatecaCertificateAuthorityConfigSubjectConfig struct {
	// subject block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate_authority#subject PrivatecaCertificateAuthority#subject}
	Subject *PrivatecaCertificateAuthorityConfigSubjectConfigSubject `field:"required" json:"subject" yaml:"subject"`
	// subject_alt_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate_authority#subject_alt_name PrivatecaCertificateAuthority#subject_alt_name}
	SubjectAltName *PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltName `field:"optional" json:"subjectAltName" yaml:"subjectAltName"`
}

