package computebackendbucket


type ComputeBackendBucketCdnPolicyCacheKeyPolicy struct {
	// Allows HTTP request headers (by name) to be used in the cache key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_backend_bucket#include_http_headers ComputeBackendBucket#include_http_headers}
	IncludeHttpHeaders *[]*string `field:"optional" json:"includeHttpHeaders" yaml:"includeHttpHeaders"`
	// Names of query string parameters to include in cache keys.
	//
	// Default parameters are always included. '&' and '=' will
	// be percent encoded and not treated as delimiters.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_backend_bucket#query_string_whitelist ComputeBackendBucket#query_string_whitelist}
	QueryStringWhitelist *[]*string `field:"optional" json:"queryStringWhitelist" yaml:"queryStringWhitelist"`
}

