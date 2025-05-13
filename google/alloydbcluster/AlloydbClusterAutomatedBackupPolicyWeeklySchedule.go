// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbcluster


type AlloydbClusterAutomatedBackupPolicyWeeklySchedule struct {
	// start_times block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/alloydb_cluster#start_times AlloydbCluster#start_times}
	StartTimes interface{} `field:"required" json:"startTimes" yaml:"startTimes"`
	// The days of the week to perform a backup.
	//
	// At least one day of the week must be provided. Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/alloydb_cluster#days_of_week AlloydbCluster#days_of_week}
	DaysOfWeek *[]*string `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
}

