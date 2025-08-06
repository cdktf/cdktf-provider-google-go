// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceshttproute


type NetworkServicesHttpRouteRulesActionCorsPolicy struct {
	// In response to a preflight request, setting this to true indicates that the actual request can include user credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/network_services_http_route#allow_credentials NetworkServicesHttpRoute#allow_credentials}
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// Specifies the content for Access-Control-Allow-Headers header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/network_services_http_route#allow_headers NetworkServicesHttpRoute#allow_headers}
	AllowHeaders *[]*string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// Specifies the content for Access-Control-Allow-Methods header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/network_services_http_route#allow_methods NetworkServicesHttpRoute#allow_methods}
	AllowMethods *[]*string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// Specifies the regular expression patterns that match allowed origins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/network_services_http_route#allow_origin_regexes NetworkServicesHttpRoute#allow_origin_regexes}
	AllowOriginRegexes *[]*string `field:"optional" json:"allowOriginRegexes" yaml:"allowOriginRegexes"`
	// Specifies the list of origins that will be allowed to do CORS requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/network_services_http_route#allow_origins NetworkServicesHttpRoute#allow_origins}
	AllowOrigins *[]*string `field:"optional" json:"allowOrigins" yaml:"allowOrigins"`
	// If true, the CORS policy is disabled.
	//
	// The default value is false, which indicates that the CORS policy is in effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/network_services_http_route#disabled NetworkServicesHttpRoute#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Specifies the content for Access-Control-Expose-Headers header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/network_services_http_route#expose_headers NetworkServicesHttpRoute#expose_headers}
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// Specifies how long result of a preflight request can be cached in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/network_services_http_route#max_age NetworkServicesHttpRoute#max_age}
	MaxAge *string `field:"optional" json:"maxAge" yaml:"maxAge"`
}

