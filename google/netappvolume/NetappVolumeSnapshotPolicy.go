// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume


type NetappVolumeSnapshotPolicy struct {
	// daily_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/netapp_volume#daily_schedule NetappVolume#daily_schedule}
	DailySchedule *NetappVolumeSnapshotPolicyDailySchedule `field:"optional" json:"dailySchedule" yaml:"dailySchedule"`
	// Enables automated snapshot creation according to defined schedule.
	//
	// Default is false.
	// To disable automatic snapshot creation you have to remove the whole snapshot_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/netapp_volume#enabled NetappVolume#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// hourly_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/netapp_volume#hourly_schedule NetappVolume#hourly_schedule}
	HourlySchedule *NetappVolumeSnapshotPolicyHourlySchedule `field:"optional" json:"hourlySchedule" yaml:"hourlySchedule"`
	// monthly_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/netapp_volume#monthly_schedule NetappVolume#monthly_schedule}
	MonthlySchedule *NetappVolumeSnapshotPolicyMonthlySchedule `field:"optional" json:"monthlySchedule" yaml:"monthlySchedule"`
	// weekly_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/netapp_volume#weekly_schedule NetappVolume#weekly_schedule}
	WeeklySchedule *NetappVolumeSnapshotPolicyWeeklySchedule `field:"optional" json:"weeklySchedule" yaml:"weeklySchedule"`
}

