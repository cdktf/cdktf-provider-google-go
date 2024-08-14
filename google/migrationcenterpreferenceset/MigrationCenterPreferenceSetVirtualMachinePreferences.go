// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package migrationcenterpreferenceset


type MigrationCenterPreferenceSetVirtualMachinePreferences struct {
	// Commitment plan to consider when calculating costs for virtual machine insights and recommendations.
	//
	// If you are unsure which value to set, a 3 year commitment plan is often a good value to start with. Possible values: 'COMMITMENT_PLAN_UNSPECIFIED', 'COMMITMENT_PLAN_NONE', 'COMMITMENT_PLAN_ONE_YEAR', 'COMMITMENT_PLAN_THREE_YEARS'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/migration_center_preference_set#commitment_plan MigrationCenterPreferenceSet#commitment_plan}
	CommitmentPlan *string `field:"optional" json:"commitmentPlan" yaml:"commitmentPlan"`
	// compute_engine_preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/migration_center_preference_set#compute_engine_preferences MigrationCenterPreferenceSet#compute_engine_preferences}
	ComputeEnginePreferences *MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences `field:"optional" json:"computeEnginePreferences" yaml:"computeEnginePreferences"`
	// region_preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/migration_center_preference_set#region_preferences MigrationCenterPreferenceSet#region_preferences}
	RegionPreferences *MigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferences `field:"optional" json:"regionPreferences" yaml:"regionPreferences"`
	// Sizing optimization strategy specifies the preferred strategy used when extrapolating usage data to calculate insights and recommendations for a virtual machine.
	//
	// If you are unsure which value to set, a moderate sizing optimization strategy is often a good value to start with. Possible values: 'SIZING_OPTIMIZATION_STRATEGY_UNSPECIFIED', 'SIZING_OPTIMIZATION_STRATEGY_SAME_AS_SOURCE', 'SIZING_OPTIMIZATION_STRATEGY_MODERATE', 'SIZING_OPTIMIZATION_STRATEGY_AGGRESSIVE'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/migration_center_preference_set#sizing_optimization_strategy MigrationCenterPreferenceSet#sizing_optimization_strategy}
	SizingOptimizationStrategy *string `field:"optional" json:"sizingOptimizationStrategy" yaml:"sizingOptimizationStrategy"`
	// sole_tenancy_preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/migration_center_preference_set#sole_tenancy_preferences MigrationCenterPreferenceSet#sole_tenancy_preferences}
	SoleTenancyPreferences *MigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferences `field:"optional" json:"soleTenancyPreferences" yaml:"soleTenancyPreferences"`
	// Target product for assets using this preference set.
	//
	// Specify either target product or business goal, but not both. Possible values: 'COMPUTE_MIGRATION_TARGET_PRODUCT_UNSPECIFIED', 'COMPUTE_MIGRATION_TARGET_PRODUCT_COMPUTE_ENGINE', 'COMPUTE_MIGRATION_TARGET_PRODUCT_VMWARE_ENGINE', 'COMPUTE_MIGRATION_TARGET_PRODUCT_SOLE_TENANCY'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/migration_center_preference_set#target_product MigrationCenterPreferenceSet#target_product}
	TargetProduct *string `field:"optional" json:"targetProduct" yaml:"targetProduct"`
	// vmware_engine_preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/migration_center_preference_set#vmware_engine_preferences MigrationCenterPreferenceSet#vmware_engine_preferences}
	VmwareEnginePreferences *MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferences `field:"optional" json:"vmwareEnginePreferences" yaml:"vmwareEnginePreferences"`
}

