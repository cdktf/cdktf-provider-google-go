package binaryauthorizationpolicy


type BinaryAuthorizationPolicyAdmissionWhitelistPatterns struct {
	// An image name pattern to whitelist, in the form 'registry/path/to/image'.
	//
	// This supports a trailing * as a
	// wildcard, but this is allowed only in text after the registry/
	// part.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/binary_authorization_policy#name_pattern BinaryAuthorizationPolicy#name_pattern}
	NamePattern *string `field:"required" json:"namePattern" yaml:"namePattern"`
}

