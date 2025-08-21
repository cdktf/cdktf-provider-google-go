// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorystoreinstance


type MemorystoreInstanceAutomatedBackupConfigFixedFrequencyScheduleStartTime struct {
	// Hours of a day in 24 hour format.
	//
	// Must be greater than or equal to 0 and typically must be less than or equal to 23.
	// An API may choose to allow the value "24:00:00" for scenarios like business closing time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/memorystore_instance#hours MemorystoreInstance#hours}
	Hours *float64 `field:"required" json:"hours" yaml:"hours"`
}

