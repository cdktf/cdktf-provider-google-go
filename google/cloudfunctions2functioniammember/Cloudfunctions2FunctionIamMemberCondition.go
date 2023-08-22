package cloudfunctions2functioniammember


type Cloudfunctions2FunctionIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions2_function_iam_member#expression Cloudfunctions2FunctionIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions2_function_iam_member#title Cloudfunctions2FunctionIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions2_function_iam_member#description Cloudfunctions2FunctionIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

