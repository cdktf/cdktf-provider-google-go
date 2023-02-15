package privatecacertificatetemplate


type PrivatecaCertificateTemplatePredefinedValuesCaOptions struct {
	// Optional.
	//
	// Refers to the "CA" X.509 extension, which is a boolean value. When this value is missing, the extension will be omitted from the CA certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_template#is_ca PrivatecaCertificateTemplate#is_ca}
	IsCa interface{} `field:"optional" json:"isCa" yaml:"isCa"`
	// Optional.
	//
	// Refers to the path length restriction X.509 extension. For a CA certificate, this value describes the depth of subordinate CA certificates that are allowed. If this value is less than 0, the request will fail. If this value is missing, the max path length will be omitted from the CA certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_template#max_issuer_path_length PrivatecaCertificateTemplate#max_issuer_path_length}
	MaxIssuerPathLength *float64 `field:"optional" json:"maxIssuerPathLength" yaml:"maxIssuerPathLength"`
}

