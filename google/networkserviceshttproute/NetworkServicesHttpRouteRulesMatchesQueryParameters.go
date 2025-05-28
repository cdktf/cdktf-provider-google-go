// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceshttproute


type NetworkServicesHttpRouteRulesMatchesQueryParameters struct {
	// The value of the query parameter must exactly match the contents of exactMatch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_services_http_route#exact_match NetworkServicesHttpRoute#exact_match}
	ExactMatch *string `field:"optional" json:"exactMatch" yaml:"exactMatch"`
	// Specifies that the QueryParameterMatcher matches if request contains query parameter, irrespective of whether the parameter has a value or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_services_http_route#present_match NetworkServicesHttpRoute#present_match}
	PresentMatch interface{} `field:"optional" json:"presentMatch" yaml:"presentMatch"`
	// The name of the query parameter to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_services_http_route#query_parameter NetworkServicesHttpRoute#query_parameter}
	QueryParameter *string `field:"optional" json:"queryParameter" yaml:"queryParameter"`
	// The value of the query parameter must match the regular expression specified by regexMatch.For regular expression grammar, please see https://github.com/google/re2/wiki/Syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_services_http_route#regex_match NetworkServicesHttpRoute#regex_match}
	RegexMatch *string `field:"optional" json:"regexMatch" yaml:"regexMatch"`
}

