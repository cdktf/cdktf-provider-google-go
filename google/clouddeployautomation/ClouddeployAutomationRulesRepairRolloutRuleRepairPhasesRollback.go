// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeployautomation


type ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRollback struct {
	// Optional. The starting phase ID for the Rollout. If unspecified, the Rollout will start in the stable phase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/clouddeploy_automation#destination_phase ClouddeployAutomation#destination_phase}
	DestinationPhase *string `field:"optional" json:"destinationPhase" yaml:"destinationPhase"`
	// Optional. If pending rollout exists on the target, the rollback operation will be aborted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/clouddeploy_automation#disable_rollback_if_rollout_pending ClouddeployAutomation#disable_rollback_if_rollout_pending}
	DisableRollbackIfRolloutPending interface{} `field:"optional" json:"disableRollbackIfRolloutPending" yaml:"disableRollbackIfRolloutPending"`
}

