package spannerdatabaseiambinding


type SpannerDatabaseIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/spanner_database_iam_binding#expression SpannerDatabaseIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/spanner_database_iam_binding#title SpannerDatabaseIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/spanner_database_iam_binding#description SpannerDatabaseIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

