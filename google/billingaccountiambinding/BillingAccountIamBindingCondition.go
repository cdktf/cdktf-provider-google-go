package billingaccountiambinding


type BillingAccountIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/billing_account_iam_binding#expression BillingAccountIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/billing_account_iam_binding#title BillingAccountIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/billing_account_iam_binding#description BillingAccountIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

