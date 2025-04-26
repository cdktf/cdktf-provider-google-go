// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeployautomation


type ClouddeployAutomationRulesTimedPromoteReleaseRule struct {
	// Required.
	//
	// ID of the rule. This id must be unique in the 'Automation' resource to which this rule belongs. The format is 'a-z{0,62}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/clouddeploy_automation#id ClouddeployAutomation#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Required. Schedule in crontab format. e.g. '0 9 * * 1' for every Monday at 9am.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/clouddeploy_automation#schedule ClouddeployAutomation#schedule}
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// Required. The time zone in IANA format IANA Time Zone Database (e.g. America/New_York).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/clouddeploy_automation#time_zone ClouddeployAutomation#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// Optional. The starting phase of the rollout created by this rule. Default to the first phase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/clouddeploy_automation#destination_phase ClouddeployAutomation#destination_phase}
	DestinationPhase *string `field:"optional" json:"destinationPhase" yaml:"destinationPhase"`
	// Optional.
	//
	// The ID of the stage in the pipeline to which this Release is deploying. If unspecified, default it to the next stage in the promotion flow. The value of this field could be one of the following:
	//   - The last segment of a target name
	//   - "@next", the next target in the promotion sequence"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/clouddeploy_automation#destination_target_id ClouddeployAutomation#destination_target_id}
	DestinationTargetId *string `field:"optional" json:"destinationTargetId" yaml:"destinationTargetId"`
}

