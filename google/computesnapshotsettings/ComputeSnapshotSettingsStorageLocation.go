// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesnapshotsettings


type ComputeSnapshotSettingsStorageLocation struct {
	// The chosen location policy Possible values: ["NEAREST_MULTI_REGION", "LOCAL_REGION", "SPECIFIC_LOCATIONS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/compute_snapshot_settings#policy ComputeSnapshotSettings#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/compute_snapshot_settings#locations ComputeSnapshotSettings#locations}
	Locations interface{} `field:"optional" json:"locations" yaml:"locations"`
}

