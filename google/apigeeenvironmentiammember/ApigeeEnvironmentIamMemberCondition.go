package apigeeenvironmentiammember


type ApigeeEnvironmentIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apigee_environment_iam_member#expression ApigeeEnvironmentIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apigee_environment_iam_member#title ApigeeEnvironmentIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apigee_environment_iam_member#description ApigeeEnvironmentIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

