// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploydeploypolicy


type ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows struct {
	// Required. The time zone in IANA format IANA Time Zone Database (e.g. America/New_York).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/clouddeploy_deploy_policy#time_zone ClouddeployDeployPolicy#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// one_time_windows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/clouddeploy_deploy_policy#one_time_windows ClouddeployDeployPolicy#one_time_windows}
	OneTimeWindows interface{} `field:"optional" json:"oneTimeWindows" yaml:"oneTimeWindows"`
	// weekly_windows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/clouddeploy_deploy_policy#weekly_windows ClouddeployDeployPolicy#weekly_windows}
	WeeklyWindows interface{} `field:"optional" json:"weeklyWindows" yaml:"weeklyWindows"`
}

