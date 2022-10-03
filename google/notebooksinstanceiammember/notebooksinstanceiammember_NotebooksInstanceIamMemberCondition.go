package notebooksinstanceiammember


type NotebooksInstanceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance_iam_member#expression NotebooksInstanceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance_iam_member#title NotebooksInstanceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance_iam_member#description NotebooksInstanceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

