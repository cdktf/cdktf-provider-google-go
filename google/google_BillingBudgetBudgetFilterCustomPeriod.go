// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type BillingBudgetBudgetFilterCustomPeriod struct {
	// start_date block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/billing_budget#start_date BillingBudget#start_date}
	StartDate *BillingBudgetBudgetFilterCustomPeriodStartDate `field:"required" json:"startDate" yaml:"startDate"`
	// end_date block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/billing_budget#end_date BillingBudget#end_date}
	EndDate *BillingBudgetBudgetFilterCustomPeriodEndDate `field:"optional" json:"endDate" yaml:"endDate"`
}

