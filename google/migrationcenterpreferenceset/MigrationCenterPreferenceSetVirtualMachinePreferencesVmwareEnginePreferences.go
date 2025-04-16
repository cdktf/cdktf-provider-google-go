// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package migrationcenterpreferenceset


type MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferences struct {
	// Commitment plan to consider when calculating costs for virtual machine insights and recommendations.
	//
	// If you are unsure which value to set, a 3 year commitment plan is often a good value to start with. Possible values: 'COMMITMENT_PLAN_UNSPECIFIED', 'ON_DEMAND', 'COMMITMENT_1_YEAR_MONTHLY_PAYMENTS', 'COMMITMENT_3_YEAR_MONTHLY_PAYMENTS', 'COMMITMENT_1_YEAR_UPFRONT_PAYMENT', 'COMMITMENT_3_YEAR_UPFRONT_PAYMENT',
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/migration_center_preference_set#commitment_plan MigrationCenterPreferenceSet#commitment_plan}
	CommitmentPlan *string `field:"optional" json:"commitmentPlan" yaml:"commitmentPlan"`
	// CPU overcommit ratio. Acceptable values are between 1.0 and 8.0, with 0.1 increment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/migration_center_preference_set#cpu_overcommit_ratio MigrationCenterPreferenceSet#cpu_overcommit_ratio}
	CpuOvercommitRatio *float64 `field:"optional" json:"cpuOvercommitRatio" yaml:"cpuOvercommitRatio"`
	// Memory overcommit ratio. Acceptable values are 1.0, 1.25, 1.5, 1.75 and 2.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/migration_center_preference_set#memory_overcommit_ratio MigrationCenterPreferenceSet#memory_overcommit_ratio}
	MemoryOvercommitRatio *float64 `field:"optional" json:"memoryOvercommitRatio" yaml:"memoryOvercommitRatio"`
	// The Deduplication and Compression ratio is based on the logical (Used Before) space required to store data before applying deduplication and compression, in relation to the physical (Used After) space required after applying deduplication and compression.
	//
	// Specifically, the ratio is the Used Before space divided by the Used After space. For example, if the Used Before space is 3 GB, but the physical Used After space is 1 GB, the deduplication and compression ratio is 3x. Acceptable values are between 1.0 and 4.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/migration_center_preference_set#storage_deduplication_compression_ratio MigrationCenterPreferenceSet#storage_deduplication_compression_ratio}
	StorageDeduplicationCompressionRatio *float64 `field:"optional" json:"storageDeduplicationCompressionRatio" yaml:"storageDeduplicationCompressionRatio"`
}

