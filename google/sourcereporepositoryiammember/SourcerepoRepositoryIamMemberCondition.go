package sourcereporepositoryiammember


type SourcerepoRepositoryIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/sourcerepo_repository_iam_member#expression SourcerepoRepositoryIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/sourcerepo_repository_iam_member#title SourcerepoRepositoryIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/sourcerepo_repository_iam_member#description SourcerepoRepositoryIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

