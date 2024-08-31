// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package billingbudget


type BillingBudgetBudgetFilterCustomPeriod struct {
	// start_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/billing_budget#start_date BillingBudget#start_date}
	StartDate *BillingBudgetBudgetFilterCustomPeriodStartDate `field:"required" json:"startDate" yaml:"startDate"`
	// end_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/billing_budget#end_date BillingBudget#end_date}
	EndDate *BillingBudgetBudgetFilterCustomPeriodEndDate `field:"optional" json:"endDate" yaml:"endDate"`
}

