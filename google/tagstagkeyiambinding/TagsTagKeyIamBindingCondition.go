package tagstagkeyiambinding


type TagsTagKeyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/tags_tag_key_iam_binding#expression TagsTagKeyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/tags_tag_key_iam_binding#title TagsTagKeyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/tags_tag_key_iam_binding#description TagsTagKeyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

