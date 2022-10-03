package spannerdatabaseiammember


type SpannerDatabaseIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/spanner_database_iam_member#expression SpannerDatabaseIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/spanner_database_iam_member#title SpannerDatabaseIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/spanner_database_iam_member#description SpannerDatabaseIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

