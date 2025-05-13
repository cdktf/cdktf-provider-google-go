// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceshttproute


type NetworkServicesHttpRouteRulesMatches struct {
	// The HTTP request path value should exactly match this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/network_services_http_route#full_path_match NetworkServicesHttpRoute#full_path_match}
	FullPathMatch *string `field:"optional" json:"fullPathMatch" yaml:"fullPathMatch"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/network_services_http_route#headers NetworkServicesHttpRoute#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Specifies if prefixMatch and fullPathMatch matches are case sensitive. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/network_services_http_route#ignore_case NetworkServicesHttpRoute#ignore_case}
	IgnoreCase interface{} `field:"optional" json:"ignoreCase" yaml:"ignoreCase"`
	// The HTTP request path value must begin with specified prefixMatch. prefixMatch must begin with a /.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/network_services_http_route#prefix_match NetworkServicesHttpRoute#prefix_match}
	PrefixMatch *string `field:"optional" json:"prefixMatch" yaml:"prefixMatch"`
	// query_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/network_services_http_route#query_parameters NetworkServicesHttpRoute#query_parameters}
	QueryParameters interface{} `field:"optional" json:"queryParameters" yaml:"queryParameters"`
	// The HTTP request path value must satisfy the regular expression specified by regexMatch after removing any query parameters and anchor supplied with the original URL.
	//
	// For regular expression grammar, please see https://github.com/google/re2/wiki/Syntax
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/network_services_http_route#regex_match NetworkServicesHttpRoute#regex_match}
	RegexMatch *string `field:"optional" json:"regexMatch" yaml:"regexMatch"`
}

