package sourcereporepositoryiambinding


type SourcerepoRepositoryIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sourcerepo_repository_iam_binding#expression SourcerepoRepositoryIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sourcerepo_repository_iam_binding#title SourcerepoRepositoryIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sourcerepo_repository_iam_binding#description SourcerepoRepositoryIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

