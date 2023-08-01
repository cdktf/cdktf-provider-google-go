package computeregionurlmap


type ComputeRegionUrlMapDefaultRouteActionRetryPolicy struct {
	// Specifies the allowed number retries. This number must be > 0. If not specified, defaults to 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_url_map#num_retries ComputeRegionUrlMap#num_retries}
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// per_try_timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_url_map#per_try_timeout ComputeRegionUrlMap#per_try_timeout}
	PerTryTimeout *ComputeRegionUrlMapDefaultRouteActionRetryPolicyPerTryTimeout `field:"optional" json:"perTryTimeout" yaml:"perTryTimeout"`
	// Specifies one or more conditions when this retry policy applies.
	//
	// Valid values are listed below. Only the following codes are supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true: cancelled, deadline-exceeded, internal, resource-exhausted, unavailable.
	// - 5xx : retry is attempted if the instance or endpoint responds with any 5xx response code, or if the instance or endpoint does not respond at all. For example, disconnects, reset, read timeout, connection failure, and refused streams.
	// - gateway-error : Similar to 5xx, but only applies to response codes 502, 503 or 504.
	// - connect-failure : a retry is attempted on failures connecting to the instance or endpoint. For example, connection timeouts.
	// - retriable-4xx : a retry is attempted if the instance or endpoint responds with a 4xx response code. The only error that you can retry is error code 409.
	// - refused-stream : a retry is attempted if the instance or endpoint resets the stream with a REFUSED_STREAM error code. This reset type indicates that it is safe to retry.
	// - cancelled : a retry is attempted if the gRPC status code in the response header is set to cancelled.
	// - deadline-exceeded : a retry is attempted if the gRPC status code in the response header is set to deadline-exceeded.
	// - internal :  a retry is attempted if the gRPC status code in the response header is set to internal.
	// - resource-exhausted : a retry is attempted if the gRPC status code in the response header is set to resource-exhausted.
	// - unavailable : a retry is attempted if the gRPC status code in the response header is set to unavailable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_url_map#retry_conditions ComputeRegionUrlMap#retry_conditions}
	RetryConditions *[]*string `field:"optional" json:"retryConditions" yaml:"retryConditions"`
}

