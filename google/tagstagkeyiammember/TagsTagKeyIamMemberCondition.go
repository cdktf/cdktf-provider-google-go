package tagstagkeyiammember


type TagsTagKeyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tags_tag_key_iam_member#expression TagsTagKeyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tags_tag_key_iam_member#title TagsTagKeyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tags_tag_key_iam_member#description TagsTagKeyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

