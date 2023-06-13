package computebackendbucket


type ComputeBackendBucketCdnPolicy struct {
	// bypass_cache_on_request_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_backend_bucket#bypass_cache_on_request_headers ComputeBackendBucket#bypass_cache_on_request_headers}
	BypassCacheOnRequestHeaders interface{} `field:"optional" json:"bypassCacheOnRequestHeaders" yaml:"bypassCacheOnRequestHeaders"`
	// cache_key_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_backend_bucket#cache_key_policy ComputeBackendBucket#cache_key_policy}
	CacheKeyPolicy *ComputeBackendBucketCdnPolicyCacheKeyPolicy `field:"optional" json:"cacheKeyPolicy" yaml:"cacheKeyPolicy"`
	// Specifies the cache setting for all responses from this backend.
	//
	// The possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL and CACHE_ALL_STATIC Possible values: ["USE_ORIGIN_HEADERS", "FORCE_CACHE_ALL", "CACHE_ALL_STATIC"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_backend_bucket#cache_mode ComputeBackendBucket#cache_mode}
	CacheMode *string `field:"optional" json:"cacheMode" yaml:"cacheMode"`
	// Specifies the maximum allowed TTL for cached content served by this origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_backend_bucket#client_ttl ComputeBackendBucket#client_ttl}
	ClientTtl *float64 `field:"optional" json:"clientTtl" yaml:"clientTtl"`
	// Specifies the default TTL for cached content served by this origin for responses that do not have an existing valid TTL (max-age or s-max-age).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_backend_bucket#default_ttl ComputeBackendBucket#default_ttl}
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// Specifies the maximum allowed TTL for cached content served by this origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_backend_bucket#max_ttl ComputeBackendBucket#max_ttl}
	MaxTtl *float64 `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_backend_bucket#negative_caching ComputeBackendBucket#negative_caching}
	NegativeCaching interface{} `field:"optional" json:"negativeCaching" yaml:"negativeCaching"`
	// negative_caching_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_backend_bucket#negative_caching_policy ComputeBackendBucket#negative_caching_policy}
	NegativeCachingPolicy interface{} `field:"optional" json:"negativeCachingPolicy" yaml:"negativeCachingPolicy"`
	// If true then Cloud CDN will combine multiple concurrent cache fill requests into a small number of requests to the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_backend_bucket#request_coalescing ComputeBackendBucket#request_coalescing}
	RequestCoalescing interface{} `field:"optional" json:"requestCoalescing" yaml:"requestCoalescing"`
	// Serve existing content from the cache (if available) when revalidating content with the origin, or when an error is encountered when refreshing the cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_backend_bucket#serve_while_stale ComputeBackendBucket#serve_while_stale}
	ServeWhileStale *float64 `field:"optional" json:"serveWhileStale" yaml:"serveWhileStale"`
	// Maximum number of seconds the response to a signed URL request will be considered fresh.
	//
	// After this time period,
	// the response will be revalidated before being served.
	// When serving responses to signed URL requests,
	// Cloud CDN will internally behave as though
	// all responses from this backend had a "Cache-Control: public,
	// max-age=[TTL]" header, regardless of any existing Cache-Control
	// header. The actual headers served in responses will not be altered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_backend_bucket#signed_url_cache_max_age_sec ComputeBackendBucket#signed_url_cache_max_age_sec}
	SignedUrlCacheMaxAgeSec *float64 `field:"optional" json:"signedUrlCacheMaxAgeSec" yaml:"signedUrlCacheMaxAgeSec"`
}

