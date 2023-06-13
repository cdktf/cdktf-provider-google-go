package privatecacertificateauthority


type PrivatecaCertificateAuthorityConfigX509ConfigCaOptions struct {
	// When true, the "CA" in Basic Constraints extension will be set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate_authority#is_ca PrivatecaCertificateAuthority#is_ca}
	IsCa interface{} `field:"required" json:"isCa" yaml:"isCa"`
	// Refers to the "path length constraint" in Basic Constraints extension.
	//
	// For a CA certificate, this value describes the depth of
	// subordinate CA certificates that are allowed. If this value is less than 0, the request will fail. Setting the value to 0
	// requires setting 'zero_max_issuer_path_length = true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate_authority#max_issuer_path_length PrivatecaCertificateAuthority#max_issuer_path_length}
	MaxIssuerPathLength *float64 `field:"optional" json:"maxIssuerPathLength" yaml:"maxIssuerPathLength"`
	// When true, the "CA" in Basic Constraints extension will be set to false.
	//
	// If both 'is_ca' and 'non_ca' are unset, the extension will be omitted from the CA certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate_authority#non_ca PrivatecaCertificateAuthority#non_ca}
	NonCa interface{} `field:"optional" json:"nonCa" yaml:"nonCa"`
	// When true, the "path length constraint" in Basic Constraints extension will be set to 0.
	//
	// If both 'max_issuer_path_length' and 'zero_max_issuer_path_length' are unset,
	// the max path length will be omitted from the CA certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/privateca_certificate_authority#zero_max_issuer_path_length PrivatecaCertificateAuthority#zero_max_issuer_path_length}
	ZeroMaxIssuerPathLength interface{} `field:"optional" json:"zeroMaxIssuerPathLength" yaml:"zeroMaxIssuerPathLength"`
}

