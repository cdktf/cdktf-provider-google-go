package privatecacertificate


type PrivatecaCertificateConfigA struct {
	// public_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate#public_key PrivatecaCertificate#public_key}
	PublicKey *PrivatecaCertificateConfigPublicKey `field:"required" json:"publicKey" yaml:"publicKey"`
	// subject_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate#subject_config PrivatecaCertificate#subject_config}
	SubjectConfig *PrivatecaCertificateConfigSubjectConfig `field:"required" json:"subjectConfig" yaml:"subjectConfig"`
	// x509_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate#x509_config PrivatecaCertificate#x509_config}
	X509Config *PrivatecaCertificateConfigX509Config `field:"required" json:"x509Config" yaml:"x509Config"`
}

