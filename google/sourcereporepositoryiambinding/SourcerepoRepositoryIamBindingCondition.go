package sourcereporepositoryiambinding


type SourcerepoRepositoryIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/sourcerepo_repository_iam_binding#expression SourcerepoRepositoryIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/sourcerepo_repository_iam_binding#title SourcerepoRepositoryIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/sourcerepo_repository_iam_binding#description SourcerepoRepositoryIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

