package computeregionbackendservice


type ComputeRegionBackendServiceCdnPolicyNegativeCachingPolicy struct {
	// The HTTP status code to define a TTL against.
	//
	// Only HTTP status codes 300, 301, 308, 404, 405, 410, 421, 451 and 501
	// can be specified as values, and you cannot specify a status code more than once.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_backend_service#code ComputeRegionBackendService#code}
	Code *float64 `field:"optional" json:"code" yaml:"code"`
}

