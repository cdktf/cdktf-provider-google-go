package binaryauthorizationattestor


type BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeys struct {
	// ASCII-armored representation of a PGP public key, as the entire output by the command 'gpg --export --armor foo@example.com' (either LF or CRLF line endings). When using this field, id should be left blank. The BinAuthz API handlers will calculate the ID and fill it in automatically. BinAuthz computes this ID as the OpenPGP RFC4880 V4 fingerprint, represented as upper-case hex. If id is provided by the caller, it will be overwritten by the API-calculated ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/binary_authorization_attestor#ascii_armored_pgp_public_key BinaryAuthorizationAttestor#ascii_armored_pgp_public_key}
	AsciiArmoredPgpPublicKey *string `field:"optional" json:"asciiArmoredPgpPublicKey" yaml:"asciiArmoredPgpPublicKey"`
	// A descriptive comment. This field may be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/binary_authorization_attestor#comment BinaryAuthorizationAttestor#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The ID of this public key.
	//
	// Signatures verified by BinAuthz
	// must include the ID of the public key that can be used to
	// verify them, and that ID must match the contents of this
	// field exactly. Additional restrictions on this field can
	// be imposed based on which public key type is encapsulated.
	// See the documentation on publicKey cases below for details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/binary_authorization_attestor#id BinaryAuthorizationAttestor#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// pkix_public_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/binary_authorization_attestor#pkix_public_key BinaryAuthorizationAttestor#pkix_public_key}
	PkixPublicKey *BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKey `field:"optional" json:"pkixPublicKey" yaml:"pkixPublicKey"`
}

