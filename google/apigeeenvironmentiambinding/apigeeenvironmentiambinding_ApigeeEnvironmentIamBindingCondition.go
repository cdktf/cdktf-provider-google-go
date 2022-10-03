package apigeeenvironmentiambinding


type ApigeeEnvironmentIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apigee_environment_iam_binding#expression ApigeeEnvironmentIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apigee_environment_iam_binding#title ApigeeEnvironmentIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apigee_environment_iam_binding#description ApigeeEnvironmentIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

