// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectivityspoke


type NetworkConnectivitySpokeLinkedVpcNetwork struct {
	// The URI of the VPC network resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/network_connectivity_spoke#uri NetworkConnectivitySpoke#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// IP ranges encompassing the subnets to be excluded from peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/network_connectivity_spoke#exclude_export_ranges NetworkConnectivitySpoke#exclude_export_ranges}
	ExcludeExportRanges *[]*string `field:"optional" json:"excludeExportRanges" yaml:"excludeExportRanges"`
	// IP ranges allowed to be included from peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/network_connectivity_spoke#include_export_ranges NetworkConnectivitySpoke#include_export_ranges}
	IncludeExportRanges *[]*string `field:"optional" json:"includeExportRanges" yaml:"includeExportRanges"`
}

