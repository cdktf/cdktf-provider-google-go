// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceshttproute


type NetworkServicesHttpRouteRulesActionRedirect struct {
	// The host that will be used in the redirect response instead of the one that was supplied in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/network_services_http_route#host_redirect NetworkServicesHttpRoute#host_redirect}
	HostRedirect *string `field:"optional" json:"hostRedirect" yaml:"hostRedirect"`
	// If set to true, the URL scheme in the redirected request is set to https.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/network_services_http_route#https_redirect NetworkServicesHttpRoute#https_redirect}
	HttpsRedirect interface{} `field:"optional" json:"httpsRedirect" yaml:"httpsRedirect"`
	// The path that will be used in the redirect response instead of the one that was supplied in the request.
	//
	// pathRedirect can not be supplied together with prefixRedirect. Supply one alone or neither. If neither is supplied, the path of the original request will be used for the redirect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/network_services_http_route#path_redirect NetworkServicesHttpRoute#path_redirect}
	PathRedirect *string `field:"optional" json:"pathRedirect" yaml:"pathRedirect"`
	// The port that will be used in the redirected request instead of the one that was supplied in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/network_services_http_route#port_redirect NetworkServicesHttpRoute#port_redirect}
	PortRedirect *float64 `field:"optional" json:"portRedirect" yaml:"portRedirect"`
	// Indicates that during redirection, the matched prefix (or path) should be swapped with this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/network_services_http_route#prefix_rewrite NetworkServicesHttpRoute#prefix_rewrite}
	PrefixRewrite *string `field:"optional" json:"prefixRewrite" yaml:"prefixRewrite"`
	// The HTTP Status code to use for the redirect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/network_services_http_route#response_code NetworkServicesHttpRoute#response_code}
	ResponseCode *string `field:"optional" json:"responseCode" yaml:"responseCode"`
	// If set to true, any accompanying query portion of the original URL is removed prior to redirecting the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/network_services_http_route#strip_query NetworkServicesHttpRoute#strip_query}
	StripQuery interface{} `field:"optional" json:"stripQuery" yaml:"stripQuery"`
}

