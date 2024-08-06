// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionurlmap


type ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderAction struct {
	// request_headers_to_add block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/compute_region_url_map#request_headers_to_add ComputeRegionUrlMap#request_headers_to_add}
	RequestHeadersToAdd interface{} `field:"optional" json:"requestHeadersToAdd" yaml:"requestHeadersToAdd"`
	// A list of header names for headers that need to be removed from the request before forwarding the request to the backendService.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/compute_region_url_map#request_headers_to_remove ComputeRegionUrlMap#request_headers_to_remove}
	RequestHeadersToRemove *[]*string `field:"optional" json:"requestHeadersToRemove" yaml:"requestHeadersToRemove"`
	// response_headers_to_add block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/compute_region_url_map#response_headers_to_add ComputeRegionUrlMap#response_headers_to_add}
	ResponseHeadersToAdd interface{} `field:"optional" json:"responseHeadersToAdd" yaml:"responseHeadersToAdd"`
	// A list of header names for headers that need to be removed from the response before sending the response back to the client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/compute_region_url_map#response_headers_to_remove ComputeRegionUrlMap#response_headers_to_remove}
	ResponseHeadersToRemove *[]*string `field:"optional" json:"responseHeadersToRemove" yaml:"responseHeadersToRemove"`
}

