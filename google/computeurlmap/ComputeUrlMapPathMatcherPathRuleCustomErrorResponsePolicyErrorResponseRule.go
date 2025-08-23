// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeurlmap


type ComputeUrlMapPathMatcherPathRuleCustomErrorResponsePolicyErrorResponseRule struct {
	// Valid values include:.
	//
	// - A number between 400 and 599: For example 401 or 503, in which case the load balancer applies the policy if the error code exactly matches this value.
	// - 5xx: Load Balancer will apply the policy if the backend service responds with any response code in the range of 500 to 599.
	// - 4xx: Load Balancer will apply the policy if the backend service responds with any response code in the range of 400 to 499.
	//
	// Values must be unique within matchResponseCodes and across all errorResponseRules of CustomErrorResponsePolicy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/compute_url_map#match_response_codes ComputeUrlMap#match_response_codes}
	MatchResponseCodes *[]*string `field:"optional" json:"matchResponseCodes" yaml:"matchResponseCodes"`
	// The HTTP status code returned with the response containing the custom error content.
	//
	// If overrideResponseCode is not supplied, the same response code returned by the original backend bucket or backend service is returned to the client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/compute_url_map#override_response_code ComputeUrlMap#override_response_code}
	OverrideResponseCode *float64 `field:"optional" json:"overrideResponseCode" yaml:"overrideResponseCode"`
	// The full path to a file within backendBucket .
	//
	// For example: /errors/defaultError.html
	// path must start with a leading slash. path cannot have trailing slashes.
	// If the file is not available in backendBucket or the load balancer cannot reach the BackendBucket, a simple Not Found Error is returned to the client.
	// The value must be from 1 to 1024 characters
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/compute_url_map#path ComputeUrlMap#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

