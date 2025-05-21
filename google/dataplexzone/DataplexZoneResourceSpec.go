// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexzone


type DataplexZoneResourceSpec struct {
	// Required.
	//
	// Immutable. The location type of the resources that are allowed to be attached to the assets within this zone. Possible values: LOCATION_TYPE_UNSPECIFIED, SINGLE_REGION, MULTI_REGION
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/dataplex_zone#location_type DataplexZone#location_type}
	LocationType *string `field:"required" json:"locationType" yaml:"locationType"`
}

