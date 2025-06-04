// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceshttproute


type NetworkServicesHttpRouteRulesMatchesHeaders struct {
	// The value of the header should match exactly the content of exactMatch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/network_services_http_route#exact_match NetworkServicesHttpRoute#exact_match}
	ExactMatch *string `field:"optional" json:"exactMatch" yaml:"exactMatch"`
	// The name of the HTTP header to match against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/network_services_http_route#header NetworkServicesHttpRoute#header}
	Header *string `field:"optional" json:"header" yaml:"header"`
	// If specified, the match result will be inverted before checking. Default value is set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/network_services_http_route#invert_match NetworkServicesHttpRoute#invert_match}
	InvertMatch interface{} `field:"optional" json:"invertMatch" yaml:"invertMatch"`
	// The value of the header must start with the contents of prefixMatch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/network_services_http_route#prefix_match NetworkServicesHttpRoute#prefix_match}
	PrefixMatch *string `field:"optional" json:"prefixMatch" yaml:"prefixMatch"`
	// A header with headerName must exist. The match takes place whether or not the header has a value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/network_services_http_route#present_match NetworkServicesHttpRoute#present_match}
	PresentMatch interface{} `field:"optional" json:"presentMatch" yaml:"presentMatch"`
	// range_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/network_services_http_route#range_match NetworkServicesHttpRoute#range_match}
	RangeMatch *NetworkServicesHttpRouteRulesMatchesHeadersRangeMatch `field:"optional" json:"rangeMatch" yaml:"rangeMatch"`
	// The value of the header must match the regular expression specified in regexMatch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/network_services_http_route#regex_match NetworkServicesHttpRoute#regex_match}
	RegexMatch *string `field:"optional" json:"regexMatch" yaml:"regexMatch"`
	// The value of the header must end with the contents of suffixMatch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/network_services_http_route#suffix_match NetworkServicesHttpRoute#suffix_match}
	SuffixMatch *string `field:"optional" json:"suffixMatch" yaml:"suffixMatch"`
}

