package apigeeenvironmentiammember


type ApigeeEnvironmentIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/apigee_environment_iam_member#expression ApigeeEnvironmentIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/apigee_environment_iam_member#title ApigeeEnvironmentIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/apigee_environment_iam_member#description ApigeeEnvironmentIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

