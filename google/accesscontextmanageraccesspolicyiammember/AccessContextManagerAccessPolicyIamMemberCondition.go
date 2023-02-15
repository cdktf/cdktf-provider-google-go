package accesscontextmanageraccesspolicyiammember


type AccessContextManagerAccessPolicyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/access_context_manager_access_policy_iam_member#expression AccessContextManagerAccessPolicyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/access_context_manager_access_policy_iam_member#title AccessContextManagerAccessPolicyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/access_context_manager_access_policy_iam_member#description AccessContextManagerAccessPolicyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

