package spannerinstanceiambinding


type SpannerInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/spanner_instance_iam_binding#expression SpannerInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/spanner_instance_iam_binding#title SpannerInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/spanner_instance_iam_binding#description SpannerInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

