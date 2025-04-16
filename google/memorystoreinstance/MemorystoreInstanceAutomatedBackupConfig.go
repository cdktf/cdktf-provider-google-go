// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorystoreinstance


type MemorystoreInstanceAutomatedBackupConfig struct {
	// fixed_frequency_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/memorystore_instance#fixed_frequency_schedule MemorystoreInstance#fixed_frequency_schedule}
	FixedFrequencySchedule *MemorystoreInstanceAutomatedBackupConfigFixedFrequencySchedule `field:"required" json:"fixedFrequencySchedule" yaml:"fixedFrequencySchedule"`
	// How long to keep automated backups before the backups are deleted.
	//
	// The value should be between 1 day and 365 days. If not specified, the default value is 35 days.
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s". The default_value is "3024000s"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/memorystore_instance#retention MemorystoreInstance#retention}
	Retention *string `field:"required" json:"retention" yaml:"retention"`
}

