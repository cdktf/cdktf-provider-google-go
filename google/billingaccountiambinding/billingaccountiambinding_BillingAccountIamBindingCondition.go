package billingaccountiambinding


type BillingAccountIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/billing_account_iam_binding#expression BillingAccountIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/billing_account_iam_binding#title BillingAccountIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/billing_account_iam_binding#description BillingAccountIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

