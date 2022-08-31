// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type PrivatecaCertificateAuthorityConfigX509ConfigPolicyIds struct {
	// An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_authority#object_id_path PrivatecaCertificateAuthority#object_id_path}
	ObjectIdPath *[]*float64 `field:"required" json:"objectIdPath" yaml:"objectIdPath"`
}

