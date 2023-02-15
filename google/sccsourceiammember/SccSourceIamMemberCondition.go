package sccsourceiammember


type SccSourceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/scc_source_iam_member#expression SccSourceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/scc_source_iam_member#title SccSourceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/scc_source_iam_member#description SccSourceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

