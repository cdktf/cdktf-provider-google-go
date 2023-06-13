package spannerdatabaseiammember


type SpannerDatabaseIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/spanner_database_iam_member#expression SpannerDatabaseIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/spanner_database_iam_member#title SpannerDatabaseIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/spanner_database_iam_member#description SpannerDatabaseIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

