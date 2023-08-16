package billingaccountiammember


type BillingAccountIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/billing_account_iam_member#expression BillingAccountIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/billing_account_iam_member#title BillingAccountIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/billing_account_iam_member#description BillingAccountIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

