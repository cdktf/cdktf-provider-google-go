package accesscontextmanageraccesspolicyiambinding


type AccessContextManagerAccessPolicyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.0/docs/resources/access_context_manager_access_policy_iam_binding#expression AccessContextManagerAccessPolicyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.0/docs/resources/access_context_manager_access_policy_iam_binding#title AccessContextManagerAccessPolicyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.0/docs/resources/access_context_manager_access_policy_iam_binding#description AccessContextManagerAccessPolicyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

