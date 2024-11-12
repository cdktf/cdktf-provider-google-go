// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package migrationcenterpreferenceset


type MigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferences struct {
	// A list of preferred regions, ordered by the most preferred region first.
	//
	// Set only valid Google Cloud region names. See https://cloud.google.com/compute/docs/regions-zones for available regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/migration_center_preference_set#preferred_regions MigrationCenterPreferenceSet#preferred_regions}
	PreferredRegions *[]*string `field:"optional" json:"preferredRegions" yaml:"preferredRegions"`
}

