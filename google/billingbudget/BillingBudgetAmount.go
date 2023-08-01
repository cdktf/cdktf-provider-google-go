package billingbudget


type BillingBudgetAmount struct {
	// Configures a budget amount that is automatically set to 100% of last period's spend.
	//
	// Boolean. Set value to true to use. Do not set to false, instead
	// use the 'specified_amount' block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/billing_budget#last_period_amount BillingBudget#last_period_amount}
	LastPeriodAmount interface{} `field:"optional" json:"lastPeriodAmount" yaml:"lastPeriodAmount"`
	// specified_amount block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/billing_budget#specified_amount BillingBudget#specified_amount}
	SpecifiedAmount *BillingBudgetAmountSpecifiedAmount `field:"optional" json:"specifiedAmount" yaml:"specifiedAmount"`
}

