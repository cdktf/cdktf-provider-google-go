package folderiambinding


type FolderIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/folder_iam_binding#expression FolderIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/folder_iam_binding#title FolderIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/folder_iam_binding#description FolderIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

