// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerydatapolicydatapolicyiambinding


type BigqueryDatapolicyDataPolicyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/bigquery_datapolicy_data_policy_iam_binding#expression BigqueryDatapolicyDataPolicyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/bigquery_datapolicy_data_policy_iam_binding#title BigqueryDatapolicyDataPolicyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/bigquery_datapolicy_data_policy_iam_binding#description BigqueryDatapolicyDataPolicyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

