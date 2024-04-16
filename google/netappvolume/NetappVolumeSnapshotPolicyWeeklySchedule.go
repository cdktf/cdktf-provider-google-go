// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume


type NetappVolumeSnapshotPolicyWeeklySchedule struct {
	// The maximum number of snapshots to keep for the weekly schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.25.0/docs/resources/netapp_volume#snapshots_to_keep NetappVolume#snapshots_to_keep}
	SnapshotsToKeep *float64 `field:"required" json:"snapshotsToKeep" yaml:"snapshotsToKeep"`
	// Set the day or days of the week to make a snapshot.
	//
	// Accepts a comma separated days of the week. Defaults to 'Sunday'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.25.0/docs/resources/netapp_volume#day NetappVolume#day}
	Day *string `field:"optional" json:"day" yaml:"day"`
	// Set the hour to create the snapshot (0-23), defaults to midnight (0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.25.0/docs/resources/netapp_volume#hour NetappVolume#hour}
	Hour *float64 `field:"optional" json:"hour" yaml:"hour"`
	// Set the minute of the hour to create the snapshot (0-59), defaults to the top of the hour (0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.25.0/docs/resources/netapp_volume#minute NetappVolume#minute}
	Minute *float64 `field:"optional" json:"minute" yaml:"minute"`
}

