package sccsourceiambinding


type SccSourceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/scc_source_iam_binding#expression SccSourceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/scc_source_iam_binding#title SccSourceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/scc_source_iam_binding#description SccSourceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

