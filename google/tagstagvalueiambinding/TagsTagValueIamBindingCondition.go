package tagstagvalueiambinding


type TagsTagValueIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/tags_tag_value_iam_binding#expression TagsTagValueIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/tags_tag_value_iam_binding#title TagsTagValueIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/tags_tag_value_iam_binding#description TagsTagValueIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

