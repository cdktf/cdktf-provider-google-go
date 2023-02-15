package projectiambinding


type ProjectIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_iam_binding#expression ProjectIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_iam_binding#title ProjectIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_iam_binding#description ProjectIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

