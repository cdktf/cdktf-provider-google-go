package privatecacertificate


type PrivatecaCertificateConfigSubjectConfig struct {
	// subject block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/privateca_certificate#subject PrivatecaCertificate#subject}
	Subject *PrivatecaCertificateConfigSubjectConfigSubject `field:"required" json:"subject" yaml:"subject"`
	// subject_alt_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/privateca_certificate#subject_alt_name PrivatecaCertificate#subject_alt_name}
	SubjectAltName *PrivatecaCertificateConfigSubjectConfigSubjectAltName `field:"optional" json:"subjectAltName" yaml:"subjectAltName"`
}

