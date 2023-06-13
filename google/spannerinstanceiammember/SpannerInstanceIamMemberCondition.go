package spannerinstanceiammember


type SpannerInstanceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/spanner_instance_iam_member#expression SpannerInstanceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/spanner_instance_iam_member#title SpannerInstanceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/spanner_instance_iam_member#description SpannerInstanceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

