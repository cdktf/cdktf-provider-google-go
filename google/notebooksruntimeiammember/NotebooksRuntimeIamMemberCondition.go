package notebooksruntimeiammember


type NotebooksRuntimeIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime_iam_member#expression NotebooksRuntimeIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime_iam_member#title NotebooksRuntimeIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime_iam_member#description NotebooksRuntimeIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

