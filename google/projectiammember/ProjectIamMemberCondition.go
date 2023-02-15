package projectiammember


type ProjectIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_iam_member#expression ProjectIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_iam_member#title ProjectIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_iam_member#description ProjectIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

