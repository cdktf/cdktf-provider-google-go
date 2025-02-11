// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume


type NetappVolumeSnapshotPolicyHourlySchedule struct {
	// The maximum number of snapshots to keep for the hourly schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_volume#snapshots_to_keep NetappVolume#snapshots_to_keep}
	SnapshotsToKeep *float64 `field:"required" json:"snapshotsToKeep" yaml:"snapshotsToKeep"`
	// Set the minute of the hour to create the snapshot (0-59), defaults to the top of the hour (0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_volume#minute NetappVolume#minute}
	Minute *float64 `field:"optional" json:"minute" yaml:"minute"`
}

