package binaryauthorizationattestor


type BinaryAuthorizationAttestorAttestationAuthorityNote struct {
	// The resource name of a ATTESTATION_AUTHORITY Note, created by the user.
	//
	// If the Note is in a different project from the Attestor, it
	// should be specified in the format 'projects/*\/notes/*' (or the legacy
	// 'providers/*\/notes/*'). This field may not be updated.
	// An attestation by this attestor is stored as a Container Analysis
	// ATTESTATION_AUTHORITY Occurrence that names a container image
	// and that links to this Note.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/binary_authorization_attestor#note_reference BinaryAuthorizationAttestor#note_reference}
	NoteReference *string `field:"required" json:"noteReference" yaml:"noteReference"`
	// public_keys block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/binary_authorization_attestor#public_keys BinaryAuthorizationAttestor#public_keys}
	PublicKeys interface{} `field:"optional" json:"publicKeys" yaml:"publicKeys"`
}

