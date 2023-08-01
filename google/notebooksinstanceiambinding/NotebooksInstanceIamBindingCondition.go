package notebooksinstanceiambinding


type NotebooksInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/notebooks_instance_iam_binding#expression NotebooksInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/notebooks_instance_iam_binding#title NotebooksInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/notebooks_instance_iam_binding#description NotebooksInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

