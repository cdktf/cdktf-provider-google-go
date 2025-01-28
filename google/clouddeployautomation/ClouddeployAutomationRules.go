// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeployautomation


type ClouddeployAutomationRules struct {
	// advance_rollout_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/clouddeploy_automation#advance_rollout_rule ClouddeployAutomation#advance_rollout_rule}
	AdvanceRolloutRule *ClouddeployAutomationRulesAdvanceRolloutRule `field:"optional" json:"advanceRolloutRule" yaml:"advanceRolloutRule"`
	// promote_release_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/clouddeploy_automation#promote_release_rule ClouddeployAutomation#promote_release_rule}
	PromoteReleaseRule *ClouddeployAutomationRulesPromoteReleaseRule `field:"optional" json:"promoteReleaseRule" yaml:"promoteReleaseRule"`
}

