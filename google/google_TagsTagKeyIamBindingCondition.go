// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type TagsTagKeyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tags_tag_key_iam_binding#expression TagsTagKeyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tags_tag_key_iam_binding#title TagsTagKeyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tags_tag_key_iam_binding#description TagsTagKeyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

