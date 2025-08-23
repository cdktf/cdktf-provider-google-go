// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume


type NetappVolumeSnapshotPolicyMonthlySchedule struct {
	// The maximum number of snapshots to keep for the monthly schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/netapp_volume#snapshots_to_keep NetappVolume#snapshots_to_keep}
	SnapshotsToKeep *float64 `field:"required" json:"snapshotsToKeep" yaml:"snapshotsToKeep"`
	// Set the day or days of the month to make a snapshot (1-31).
	//
	// Accepts a comma separated number of days. Defaults to '1'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/netapp_volume#days_of_month NetappVolume#days_of_month}
	DaysOfMonth *string `field:"optional" json:"daysOfMonth" yaml:"daysOfMonth"`
	// Set the hour to create the snapshot (0-23), defaults to midnight (0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/netapp_volume#hour NetappVolume#hour}
	Hour *float64 `field:"optional" json:"hour" yaml:"hour"`
	// Set the minute of the hour to create the snapshot (0-59), defaults to the top of the hour (0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/netapp_volume#minute NetappVolume#minute}
	Minute *float64 `field:"optional" json:"minute" yaml:"minute"`
}

