// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeurlmap


type ComputeUrlMapPathMatcherRouteRulesUrlRedirect struct {
	// The host that will be used in the redirect response instead of the one that was supplied in the request.
	//
	// The value must be between 1 and 255 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/compute_url_map#host_redirect ComputeUrlMap#host_redirect}
	HostRedirect *string `field:"optional" json:"hostRedirect" yaml:"hostRedirect"`
	// If set to true, the URL scheme in the redirected request is set to https.
	//
	// If set
	// to false, the URL scheme of the redirected request will remain the same as that
	// of the request. This must only be set for UrlMaps used in TargetHttpProxys.
	// Setting this true for TargetHttpsProxy is not permitted. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/compute_url_map#https_redirect ComputeUrlMap#https_redirect}
	HttpsRedirect interface{} `field:"optional" json:"httpsRedirect" yaml:"httpsRedirect"`
	// The path that will be used in the redirect response instead of the one that was supplied in the request.
	//
	// Only one of pathRedirect or prefixRedirect must be
	// specified. The value must be between 1 and 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/compute_url_map#path_redirect ComputeUrlMap#path_redirect}
	PathRedirect *string `field:"optional" json:"pathRedirect" yaml:"pathRedirect"`
	// The prefix that replaces the prefixMatch specified in the HttpRouteRuleMatch, retaining the remaining portion of the URL before redirecting the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/compute_url_map#prefix_redirect ComputeUrlMap#prefix_redirect}
	PrefixRedirect *string `field:"optional" json:"prefixRedirect" yaml:"prefixRedirect"`
	// The HTTP Status code to use for this RedirectAction. Supported values are:.
	//
	// * MOVED_PERMANENTLY_DEFAULT, which is the default value and corresponds to 301.
	//
	// * FOUND, which corresponds to 302.
	//
	// * SEE_OTHER which corresponds to 303.
	//
	// * TEMPORARY_REDIRECT, which corresponds to 307. In this case, the request method will be retained.
	//
	// * PERMANENT_REDIRECT, which corresponds to 308. In this case, the request method will be retained. Possible values: ["FOUND", "MOVED_PERMANENTLY_DEFAULT", "PERMANENT_REDIRECT", "SEE_OTHER", "TEMPORARY_REDIRECT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/compute_url_map#redirect_response_code ComputeUrlMap#redirect_response_code}
	RedirectResponseCode *string `field:"optional" json:"redirectResponseCode" yaml:"redirectResponseCode"`
	// If set to true, any accompanying query portion of the original URL is removed prior to redirecting the request.
	//
	// If set to false, the query portion of the
	// original URL is retained. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/compute_url_map#strip_query ComputeUrlMap#strip_query}
	StripQuery interface{} `field:"optional" json:"stripQuery" yaml:"stripQuery"`
}

