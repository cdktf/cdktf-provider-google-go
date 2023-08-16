package networkservicesedgecachekeyset


type NetworkServicesEdgeCacheKeysetValidationSharedKeys struct {
	// The name of the secret version in Secret Manager.
	//
	// The resource name of the secret version must be in the format 'projects/*\/secrets/*\/versions/*' where the '*' values are replaced by the secrets themselves.
	// The secrets must be at least 16 bytes large.  The recommended secret size depends on the signature algorithm you are using.
	// If you are using HMAC-SHA1, we suggest 20-byte secrets.
	// If you are using HMAC-SHA256, we suggest 32-byte secrets.
	// See RFC 2104, Section 3 for more details on these recommendations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_keyset#secret_version NetworkServicesEdgeCacheKeyset#secret_version}
	SecretVersion *string `field:"required" json:"secretVersion" yaml:"secretVersion"`
}

