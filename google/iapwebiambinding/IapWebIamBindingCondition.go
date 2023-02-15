package iapwebiambinding


type IapWebIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_web_iam_binding#expression IapWebIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_web_iam_binding#title IapWebIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_web_iam_binding#description IapWebIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

