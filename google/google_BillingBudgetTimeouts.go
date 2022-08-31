// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type BillingBudgetTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/billing_budget#create BillingBudget#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/billing_budget#delete BillingBudget#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/billing_budget#update BillingBudget#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

