// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeployautomation


type ClouddeployAutomationRulesAdvanceRolloutRule struct {
	// Required.
	//
	// ID of the rule. This id must be unique in the 'Automation' resource to which this rule belongs. The format is 'a-z{0,62}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/clouddeploy_automation#id ClouddeployAutomation#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Optional.
	//
	// Proceeds only after phase name matched any one in the list. This value must consist of lower-case letters, numbers, and hyphens, start with a letter and end with a letter or a number, and have a max length of 63 characters. In other words, it must match the following regex: '^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/clouddeploy_automation#source_phases ClouddeployAutomation#source_phases}
	SourcePhases *[]*string `field:"optional" json:"sourcePhases" yaml:"sourcePhases"`
	// Optional. How long to wait after a rollout is finished.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/clouddeploy_automation#wait ClouddeployAutomation#wait}
	Wait *string `field:"optional" json:"wait" yaml:"wait"`
}

