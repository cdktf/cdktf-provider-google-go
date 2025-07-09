// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesnapshotsettings


type ComputeSnapshotSettingsStorageLocationLocations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/compute_snapshot_settings#location ComputeSnapshotSettings#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Name of the location.
	//
	// It should be one of the Cloud Storage buckets.
	// Only one location can be specified. (should match location)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/compute_snapshot_settings#name ComputeSnapshotSettings#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

