// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesedgecacheorigin


type NetworkServicesEdgeCacheOriginFlexShielding struct {
	// Whenever possible, content will be fetched from origin and cached in or near the specified origin. Best effort.
	//
	// You must specify exactly one FlexShieldingRegion. Possible values: ["AFRICA_SOUTH1", "ME_CENTRAL1"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/network_services_edge_cache_origin#flex_shielding_regions NetworkServicesEdgeCacheOrigin#flex_shielding_regions}
	FlexShieldingRegions *[]*string `field:"optional" json:"flexShieldingRegions" yaml:"flexShieldingRegions"`
}

