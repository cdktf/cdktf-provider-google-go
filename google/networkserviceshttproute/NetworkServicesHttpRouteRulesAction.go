// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceshttproute


type NetworkServicesHttpRouteRulesAction struct {
	// cors_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_http_route#cors_policy NetworkServicesHttpRoute#cors_policy}
	CorsPolicy *NetworkServicesHttpRouteRulesActionCorsPolicy `field:"optional" json:"corsPolicy" yaml:"corsPolicy"`
	// destinations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_http_route#destinations NetworkServicesHttpRoute#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// fault_injection_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_http_route#fault_injection_policy NetworkServicesHttpRoute#fault_injection_policy}
	FaultInjectionPolicy *NetworkServicesHttpRouteRulesActionFaultInjectionPolicy `field:"optional" json:"faultInjectionPolicy" yaml:"faultInjectionPolicy"`
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_http_route#redirect NetworkServicesHttpRoute#redirect}
	Redirect *NetworkServicesHttpRouteRulesActionRedirect `field:"optional" json:"redirect" yaml:"redirect"`
	// request_header_modifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_http_route#request_header_modifier NetworkServicesHttpRoute#request_header_modifier}
	RequestHeaderModifier *NetworkServicesHttpRouteRulesActionRequestHeaderModifier `field:"optional" json:"requestHeaderModifier" yaml:"requestHeaderModifier"`
	// request_mirror_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_http_route#request_mirror_policy NetworkServicesHttpRoute#request_mirror_policy}
	RequestMirrorPolicy *NetworkServicesHttpRouteRulesActionRequestMirrorPolicy `field:"optional" json:"requestMirrorPolicy" yaml:"requestMirrorPolicy"`
	// response_header_modifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_http_route#response_header_modifier NetworkServicesHttpRoute#response_header_modifier}
	ResponseHeaderModifier *NetworkServicesHttpRouteRulesActionResponseHeaderModifier `field:"optional" json:"responseHeaderModifier" yaml:"responseHeaderModifier"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_http_route#retry_policy NetworkServicesHttpRoute#retry_policy}
	RetryPolicy *NetworkServicesHttpRouteRulesActionRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// Specifies the timeout for selected route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_http_route#timeout NetworkServicesHttpRoute#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	// url_rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_http_route#url_rewrite NetworkServicesHttpRoute#url_rewrite}
	UrlRewrite *NetworkServicesHttpRouteRulesActionUrlRewrite `field:"optional" json:"urlRewrite" yaml:"urlRewrite"`
}

