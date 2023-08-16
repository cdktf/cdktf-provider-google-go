package billingbudget


type BillingBudgetThresholdRules struct {
	// Send an alert when this threshold is exceeded.
	//
	// This is a
	// 1.0-based percentage, so 0.5 = 50%. Must be >= 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/billing_budget#threshold_percent BillingBudget#threshold_percent}
	ThresholdPercent *float64 `field:"required" json:"thresholdPercent" yaml:"thresholdPercent"`
	// The type of basis used to determine if spend has passed the threshold. Default value: "CURRENT_SPEND" Possible values: ["CURRENT_SPEND", "FORECASTED_SPEND"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/billing_budget#spend_basis BillingBudget#spend_basis}
	SpendBasis *string `field:"optional" json:"spendBasis" yaml:"spendBasis"`
}

