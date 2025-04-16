// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceshttproute


type NetworkServicesHttpRouteRulesMatchesHeadersRangeMatch struct {
	// End of the range (exclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/network_services_http_route#end NetworkServicesHttpRoute#end}
	End *float64 `field:"required" json:"end" yaml:"end"`
	// Start of the range (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/network_services_http_route#start NetworkServicesHttpRoute#start}
	Start *float64 `field:"required" json:"start" yaml:"start"`
}

