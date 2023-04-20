package tagstagvalueiammember


type TagsTagValueIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.1/docs/resources/tags_tag_value_iam_member#expression TagsTagValueIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.1/docs/resources/tags_tag_value_iam_member#title TagsTagValueIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.1/docs/resources/tags_tag_value_iam_member#description TagsTagValueIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

