// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type NotebooksRuntimeIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_runtime_iam_member#expression NotebooksRuntimeIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_runtime_iam_member#title NotebooksRuntimeIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_runtime_iam_member#description NotebooksRuntimeIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

