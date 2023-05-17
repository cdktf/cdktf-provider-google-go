package serviceaccountiambinding


type ServiceAccountIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/service_account_iam_binding#expression ServiceAccountIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/service_account_iam_binding#title ServiceAccountIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/service_account_iam_binding#description ServiceAccountIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

