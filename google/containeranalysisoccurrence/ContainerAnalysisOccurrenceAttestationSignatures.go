package containeranalysisoccurrence


type ContainerAnalysisOccurrenceAttestationSignatures struct {
	// The identifier for the public key that verifies this signature.
	//
	// MUST be an RFC3986 conformant
	// URI. * When possible, the key id should be an
	// immutable reference, such as a cryptographic digest.
	// Examples of valid values:
	//
	// OpenPGP V4 public key fingerprint. See https://www.iana.org/assignments/uri-schemes/prov/openpgp4fpr
	// for more details on this scheme.
	//  'openpgp4fpr:74FAF3B861BDA0870C7B6DEF607E48D2A663AEEA'
	// RFC6920 digest-named SubjectPublicKeyInfo (digest of the DER serialization):
	//  "ni:///sha-256;cD9o9Cq6LG3jD0iKXqEi_vdjJGecm_iXkbqVoScViaU"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.64.0/docs/resources/container_analysis_occurrence#public_key_id ContainerAnalysisOccurrence#public_key_id}
	PublicKeyId *string `field:"required" json:"publicKeyId" yaml:"publicKeyId"`
	// The content of the signature, an opaque bytestring.
	//
	// The payload that this signature verifies MUST be
	// unambiguously provided with the Signature during
	// verification. A wrapper message might provide the
	// payload explicitly. Alternatively, a message might
	// have a canonical serialization that can always be
	// unambiguously computed to derive the payload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.64.0/docs/resources/container_analysis_occurrence#signature ContainerAnalysisOccurrence#signature}
	Signature *string `field:"optional" json:"signature" yaml:"signature"`
}

