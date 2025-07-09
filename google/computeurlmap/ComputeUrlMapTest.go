// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeurlmap


type ComputeUrlMapTest struct {
	// Host portion of the URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/compute_url_map#host ComputeUrlMap#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Path portion of the URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/compute_url_map#path ComputeUrlMap#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Description of this test case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/compute_url_map#description ComputeUrlMap#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The expected output URL evaluated by the load balancer containing the scheme, host, path and query parameters.
	//
	// For rules that forward requests to backends, the test passes only when expectedOutputUrl matches the request forwarded by the load balancer to backends. For rules with urlRewrite, the test verifies that the forwarded request matches hostRewrite and pathPrefixRewrite in the urlRewrite action. When service is specified, expectedOutputUrl's scheme is ignored.
	//
	// For rules with urlRedirect, the test passes only if expectedOutputUrl matches the URL in the load balancer's redirect response. If urlRedirect specifies httpsRedirect, the test passes only if the scheme in expectedOutputUrl is also set to HTTPS. If urlRedirect specifies stripQuery, the test passes only if expectedOutputUrl does not contain any query parameters.
	//
	// expectedOutputUrl is optional when service is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/compute_url_map#expected_output_url ComputeUrlMap#expected_output_url}
	ExpectedOutputUrl *string `field:"optional" json:"expectedOutputUrl" yaml:"expectedOutputUrl"`
	// For rules with urlRedirect, the test passes only if expectedRedirectResponseCode matches the HTTP status code in load balancer's redirect response.
	//
	// expectedRedirectResponseCode cannot be set when service is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/compute_url_map#expected_redirect_response_code ComputeUrlMap#expected_redirect_response_code}
	ExpectedRedirectResponseCode *float64 `field:"optional" json:"expectedRedirectResponseCode" yaml:"expectedRedirectResponseCode"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/compute_url_map#headers ComputeUrlMap#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// The backend service or backend bucket link that should be matched by this test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/compute_url_map#service ComputeUrlMap#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

