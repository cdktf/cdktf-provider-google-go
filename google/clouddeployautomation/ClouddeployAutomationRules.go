// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeployautomation


type ClouddeployAutomationRules struct {
	// advance_rollout_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/clouddeploy_automation#advance_rollout_rule ClouddeployAutomation#advance_rollout_rule}
	AdvanceRolloutRule *ClouddeployAutomationRulesAdvanceRolloutRule `field:"optional" json:"advanceRolloutRule" yaml:"advanceRolloutRule"`
	// promote_release_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/clouddeploy_automation#promote_release_rule ClouddeployAutomation#promote_release_rule}
	PromoteReleaseRule *ClouddeployAutomationRulesPromoteReleaseRule `field:"optional" json:"promoteReleaseRule" yaml:"promoteReleaseRule"`
	// repair_rollout_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/clouddeploy_automation#repair_rollout_rule ClouddeployAutomation#repair_rollout_rule}
	RepairRolloutRule *ClouddeployAutomationRulesRepairRolloutRule `field:"optional" json:"repairRolloutRule" yaml:"repairRolloutRule"`
	// timed_promote_release_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/clouddeploy_automation#timed_promote_release_rule ClouddeployAutomation#timed_promote_release_rule}
	TimedPromoteReleaseRule *ClouddeployAutomationRulesTimedPromoteReleaseRule `field:"optional" json:"timedPromoteReleaseRule" yaml:"timedPromoteReleaseRule"`
}

