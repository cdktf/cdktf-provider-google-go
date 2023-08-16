package privatecacertificatetemplate


type PrivatecaCertificateTemplatePredefinedValuesPolicyIds struct {
	// Required. The parts of an OID path. The most significant parts of the path come first.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_certificate_template#object_id_path PrivatecaCertificateTemplate#object_id_path}
	ObjectIdPath *[]*float64 `field:"required" json:"objectIdPath" yaml:"objectIdPath"`
}

