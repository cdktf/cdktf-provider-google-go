// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorystoreinstance


type MemorystoreInstanceAutomatedBackupConfigFixedFrequencySchedule struct {
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/memorystore_instance#start_time MemorystoreInstance#start_time}
	StartTime *MemorystoreInstanceAutomatedBackupConfigFixedFrequencyScheduleStartTime `field:"required" json:"startTime" yaml:"startTime"`
}

