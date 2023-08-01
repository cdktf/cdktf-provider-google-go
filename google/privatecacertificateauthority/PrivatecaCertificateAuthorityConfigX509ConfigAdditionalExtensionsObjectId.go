package privatecacertificateauthority


type PrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId struct {
	// An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate_authority#object_id_path PrivatecaCertificateAuthority#object_id_path}
	ObjectIdPath *[]*float64 `field:"required" json:"objectIdPath" yaml:"objectIdPath"`
}

