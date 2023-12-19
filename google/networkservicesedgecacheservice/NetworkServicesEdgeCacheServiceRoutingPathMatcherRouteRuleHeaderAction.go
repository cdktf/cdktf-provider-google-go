// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesedgecacheservice


type NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderAction struct {
	// request_header_to_add block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/network_services_edge_cache_service#request_header_to_add NetworkServicesEdgeCacheService#request_header_to_add}
	RequestHeaderToAdd interface{} `field:"optional" json:"requestHeaderToAdd" yaml:"requestHeaderToAdd"`
	// request_header_to_remove block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/network_services_edge_cache_service#request_header_to_remove NetworkServicesEdgeCacheService#request_header_to_remove}
	RequestHeaderToRemove interface{} `field:"optional" json:"requestHeaderToRemove" yaml:"requestHeaderToRemove"`
	// response_header_to_add block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/network_services_edge_cache_service#response_header_to_add NetworkServicesEdgeCacheService#response_header_to_add}
	ResponseHeaderToAdd interface{} `field:"optional" json:"responseHeaderToAdd" yaml:"responseHeaderToAdd"`
	// response_header_to_remove block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/network_services_edge_cache_service#response_header_to_remove NetworkServicesEdgeCacheService#response_header_to_remove}
	ResponseHeaderToRemove interface{} `field:"optional" json:"responseHeaderToRemove" yaml:"responseHeaderToRemove"`
}

