// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package billingbudget


type BillingBudgetBudgetFilterCustomPeriodStartDate struct {
	// Day of a month. Must be from 1 to 31 and valid for the year and month.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.84.0/docs/resources/billing_budget#day BillingBudget#day}
	Day *float64 `field:"required" json:"day" yaml:"day"`
	// Month of a year. Must be from 1 to 12.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.84.0/docs/resources/billing_budget#month BillingBudget#month}
	Month *float64 `field:"required" json:"month" yaml:"month"`
	// Year of the date. Must be from 1 to 9999.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.84.0/docs/resources/billing_budget#year BillingBudget#year}
	Year *float64 `field:"required" json:"year" yaml:"year"`
}

