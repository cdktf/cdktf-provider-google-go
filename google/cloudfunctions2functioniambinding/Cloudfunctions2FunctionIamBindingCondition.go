package cloudfunctions2functioniambinding


type Cloudfunctions2FunctionIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions2_function_iam_binding#expression Cloudfunctions2FunctionIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions2_function_iam_binding#title Cloudfunctions2FunctionIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions2_function_iam_binding#description Cloudfunctions2FunctionIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

