package networkservicesedgecachekeyset


type NetworkServicesEdgeCacheKeysetPublicKey struct {
	// The ID of the public key.
	//
	// The ID must be 1-63 characters long, and comply with RFC1035.
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]*
	// which means the first character must be a letter, and all following characters must be a dash, underscore, letter or digit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/network_services_edge_cache_keyset#id NetworkServicesEdgeCacheKeyset#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The base64-encoded value of the Ed25519 public key.
	//
	// The base64 encoding can be padded (44 bytes) or unpadded (43 bytes).
	// Representations or encodings of the public key other than this will be rejected with an error.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/network_services_edge_cache_keyset#value NetworkServicesEdgeCacheKeyset#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

