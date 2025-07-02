// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploydeploypolicy


type ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindows struct {
	// Optional.
	//
	// Days of week. If left empty, all days of the week will be included. Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/clouddeploy_deploy_policy#days_of_week ClouddeployDeployPolicy#days_of_week}
	DaysOfWeek *[]*string `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// end_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/clouddeploy_deploy_policy#end_time ClouddeployDeployPolicy#end_time}
	EndTime *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTime `field:"optional" json:"endTime" yaml:"endTime"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/clouddeploy_deploy_policy#start_time ClouddeployDeployPolicy#start_time}
	StartTime *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsStartTime `field:"optional" json:"startTime" yaml:"startTime"`
}

