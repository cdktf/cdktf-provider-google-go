package privatecacertificateauthority


type PrivatecaCertificateAuthorityConfigA struct {
	// subject_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#subject_config PrivatecaCertificateAuthority#subject_config}
	SubjectConfig *PrivatecaCertificateAuthorityConfigSubjectConfig `field:"required" json:"subjectConfig" yaml:"subjectConfig"`
	// x509_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#x509_config PrivatecaCertificateAuthority#x509_config}
	X509Config *PrivatecaCertificateAuthorityConfigX509Config `field:"required" json:"x509Config" yaml:"x509Config"`
}

