// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploydeploypolicy


type ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindows struct {
	// end_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/clouddeploy_deploy_policy#end_date ClouddeployDeployPolicy#end_date}
	EndDate *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDate `field:"required" json:"endDate" yaml:"endDate"`
	// end_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/clouddeploy_deploy_policy#end_time ClouddeployDeployPolicy#end_time}
	EndTime *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTime `field:"required" json:"endTime" yaml:"endTime"`
	// start_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/clouddeploy_deploy_policy#start_date ClouddeployDeployPolicy#start_date}
	StartDate *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDate `field:"required" json:"startDate" yaml:"startDate"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/clouddeploy_deploy_policy#start_time ClouddeployDeployPolicy#start_time}
	StartTime *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTime `field:"required" json:"startTime" yaml:"startTime"`
}

