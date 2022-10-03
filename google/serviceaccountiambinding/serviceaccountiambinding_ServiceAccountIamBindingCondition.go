package serviceaccountiambinding


type ServiceAccountIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/service_account_iam_binding#expression ServiceAccountIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/service_account_iam_binding#title ServiceAccountIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/service_account_iam_binding#description ServiceAccountIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

