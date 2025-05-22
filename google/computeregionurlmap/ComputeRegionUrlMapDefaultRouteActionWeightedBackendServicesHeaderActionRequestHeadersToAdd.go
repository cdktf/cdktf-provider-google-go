// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionurlmap


type ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAdd struct {
	// The name of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/compute_region_url_map#header_name ComputeRegionUrlMap#header_name}
	HeaderName *string `field:"optional" json:"headerName" yaml:"headerName"`
	// The value of the header to add.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/compute_region_url_map#header_value ComputeRegionUrlMap#header_value}
	HeaderValue *string `field:"optional" json:"headerValue" yaml:"headerValue"`
	// If false, headerValue is appended to any values that already exist for the header.
	//
	// If true, headerValue is set for the header, discarding any values that were set for that header.
	// The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/compute_region_url_map#replace ComputeRegionUrlMap#replace}
	Replace interface{} `field:"optional" json:"replace" yaml:"replace"`
}

