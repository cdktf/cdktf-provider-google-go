// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploydeploypolicy


type ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDate struct {
	// Day of a month. Must be from 1 to 31 and valid for the year and month.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/clouddeploy_deploy_policy#day ClouddeployDeployPolicy#day}
	Day *float64 `field:"optional" json:"day" yaml:"day"`
	// Month of a year. Must be from 1 to 12.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/clouddeploy_deploy_policy#month ClouddeployDeployPolicy#month}
	Month *float64 `field:"optional" json:"month" yaml:"month"`
	// Year of the date. Must be from 1 to 9999.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/clouddeploy_deploy_policy#year ClouddeployDeployPolicy#year}
	Year *float64 `field:"optional" json:"year" yaml:"year"`
}

