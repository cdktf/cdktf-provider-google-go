// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeployautomation


type ClouddeployAutomationRulesPromoteReleaseRule struct {
	// Required.
	//
	// ID of the rule. This id must be unique in the 'Automation' resource to which this rule belongs. The format is 'a-z{0,62}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/clouddeploy_automation#id ClouddeployAutomation#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Optional. The starting phase of the rollout created by this operation. Default to the first phase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/clouddeploy_automation#destination_phase ClouddeployAutomation#destination_phase}
	DestinationPhase *string `field:"optional" json:"destinationPhase" yaml:"destinationPhase"`
	// Optional.
	//
	// The ID of the stage in the pipeline to which this 'Release' is deploying. If unspecified, default it to the next stage in the promotion flow. The value of this field could be one of the following: * The last segment of a target name. It only needs the ID to determine if the target is one of the stages in the promotion sequence defined in the pipeline. * "@next", the next target in the promotion sequence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/clouddeploy_automation#destination_target_id ClouddeployAutomation#destination_target_id}
	DestinationTargetId *string `field:"optional" json:"destinationTargetId" yaml:"destinationTargetId"`
	// Optional. How long the release need to be paused until being promoted to the next target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/clouddeploy_automation#wait ClouddeployAutomation#wait}
	Wait *string `field:"optional" json:"wait" yaml:"wait"`
}

