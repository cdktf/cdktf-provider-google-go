package binaryauthorizationattestor


type BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKey struct {
	// A PEM-encoded public key, as described in 'https://tools.ietf.org/html/rfc7468#section-13'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/binary_authorization_attestor#public_key_pem BinaryAuthorizationAttestor#public_key_pem}
	PublicKeyPem *string `field:"optional" json:"publicKeyPem" yaml:"publicKeyPem"`
	// The signature algorithm used to verify a message against a signature using this key.
	//
	// These signature algorithm must
	// match the structure and any object identifiers encoded in
	// publicKeyPem (i.e. this algorithm must match that of the
	// public key).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/binary_authorization_attestor#signature_algorithm BinaryAuthorizationAttestor#signature_algorithm}
	SignatureAlgorithm *string `field:"optional" json:"signatureAlgorithm" yaml:"signatureAlgorithm"`
}

