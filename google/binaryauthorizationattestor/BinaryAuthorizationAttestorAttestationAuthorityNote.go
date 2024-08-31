// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package binaryauthorizationattestor


type BinaryAuthorizationAttestorAttestationAuthorityNote struct {
	// The resource name of a ATTESTATION_AUTHORITY Note, created by the user.
	//
	// If the Note is in a different project from the Attestor, it
	// should be specified in the format 'projects/* /notes/*' (or the legacy
	// 'providers/* /notes/*'). This field may not be updated.
	// An attestation by this attestor is stored as a Container Analysis
	// ATTESTATION_AUTHORITY Occurrence that names a container image
	// and that links to this Note.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/binary_authorization_attestor#note_reference BinaryAuthorizationAttestor#note_reference}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	NoteReference *string `field:"required" json:"noteReference" yaml:"noteReference"`
	// public_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/binary_authorization_attestor#public_keys BinaryAuthorizationAttestor#public_keys}
	PublicKeys interface{} `field:"optional" json:"publicKeys" yaml:"publicKeys"`
}

