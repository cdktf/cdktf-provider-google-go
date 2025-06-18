// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesedgecacheorigin


type NetworkServicesEdgeCacheOriginOriginOverrideAction struct {
	// header_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/network_services_edge_cache_origin#header_action NetworkServicesEdgeCacheOrigin#header_action}
	HeaderAction *NetworkServicesEdgeCacheOriginOriginOverrideActionHeaderAction `field:"optional" json:"headerAction" yaml:"headerAction"`
	// url_rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/network_services_edge_cache_origin#url_rewrite NetworkServicesEdgeCacheOrigin#url_rewrite}
	UrlRewrite *NetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewrite `field:"optional" json:"urlRewrite" yaml:"urlRewrite"`
}

